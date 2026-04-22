package main

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite", "../data/events.db")
	if err != nil {
		log.Fatal("DB open error:", err)
	}

	createTable := `
    CREATE TABLE IF NOT EXISTS events (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        class TEXT,
        confidence REAL,
        x INTEGER,
        y INTEGER,
        w INTEGER,
        h INTEGER,
        timestamp INTEGER
    );
    `
	if _, err := db.Exec(createTable); err != nil {
		log.Fatal("DB table create error:", err)
	}

	return db
}
