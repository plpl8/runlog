package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RunRepository struct {
	db *pgxpool.Pool
}

func NewRunRepository(db *pgxpool.Pool) *RunRepository {
	return &RunRepository{
		db: db,
	}
}

func (r *RunRepository) GetAll(ctx context.Context) ([]Run, error) {
	rows, err := r.db.Query(
		ctx,
		`
		SELECT id, date, distance, duration_seconds, type, notes, created_at
		FROM runs
		ORDER BY date DESC
		`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	runs := make([]Run, 0)

	for rows.Next() {
		var run Run

		err := rows.Scan(
			&run.ID,
			&run.Date,
			&run.Distance,
			&run.DurationSeconds,
			&run.Type,
			&run.Notes,
			&run.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		runs = append(runs, run)
	}

	return runs, nil
}
