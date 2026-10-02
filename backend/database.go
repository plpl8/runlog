package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"github.com/joho/godotenv"
)

func ConnectDB() (*pgxpool.Pool, error) {

	_ = godotenv.Load()

	connString := os.Getenv("DATABASE_URL")

	db, err := pgxpool.New(
		context.Background(),
		connString,
	)

	if err != nil {
		return nil, err
	}

	err = db.Ping(context.Background())

	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to PostgreSQL")

	return db, nil
}

