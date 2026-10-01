package main

import "time"

type Run struct {
	ID int64 `json:"id"`
	Date time.Time `json:"date"`
	Distance float64 `json:"distance"`
	DurationSeconds int `json:"duration_seconds"`
	Type string `json:"type"`
	Notes string `json:"notes"`
	CreatedAt time.Time `json:"created_at"`
}
