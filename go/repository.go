package main

import (
	"database/sql"
)

type EventRepository struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (r *EventRepository) Insert(e DetectionEvent) error {
	_, err := r.db.Exec(`
        INSERT INTO events (class, confidence, x, y, w, h, timestamp)
        VALUES (?, ?, ?, ?, ?, ?, ?)
    `,
		e.Class,
		e.Confidence,
		e.X,
		e.Y,
		e.W,
		e.H,
		e.Timestamp,
	)
	return err
}
