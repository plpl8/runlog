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

func (r *RunRepository) Create(ctx context.Context, run *Run) error {
	err := r.db.QueryRow(
		ctx,
		`
		INSERT INTO runs
		(date, distance, duration_seconds, type, notes)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at
		`,
		run.Date,
		run.Distance,
		run.DurationSeconds,
		run.Type,
		run.Notes,
	).Scan(
		&run.ID,
		&run.CreatedAt,
	)

	return err
}

func (r *RunRepository) GetByID(ctx context.Context, id int64) (*Run, error) {
	var run Run

	err := r.db.QueryRow(
		ctx,
		`
		SELECT id, date, distance, duration_seconds, type, notes, created_at
		FROM runs
		WHERE id = $1
		`,
		id,
	).Scan(
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

	return &run, nil
}

func (r *RunRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.db.Exec(
		ctx,
		`
		DELETE FROM runs
		WHERE id = $1
		`,
		id,
	)

	return err
}

func (r *RunRepository) Update(ctx context.Context, id int64, run *Run) error {
	_, err := r.db.Exec(
		ctx,
		`
		UPDATE runs
		SET date = $1,
		    distance = $2,
		    duration_seconds = $3,
		    type = $4,
		    notes = $5
		WHERE id = $6
		`,
		run.Date,
		run.Distance,
		run.DurationSeconds,
		run.Type,
		run.Notes,
		id,
	)

	return err
}
