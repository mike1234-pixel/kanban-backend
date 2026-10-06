package store

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq" // Driver side-effect import to register PostgreSQL with database/sql
)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore(connStr string) (*PostgresStore, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25) // cap active connections to prevent crashing Postgres

	db.SetMaxIdleConns(10) // keep warm connections ready so incoming requests don't wait for TCP handshakes

	db.SetConnMaxLifetime(5 * time.Minute) // retire old connections to clear stale state and handle DB restarts gracefully

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping postgres: %w", err)
	}

	return &PostgresStore{db: db}, nil
}
