package models

import "time"

type Board struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Columns   []Column  `json:"columns"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type Column struct {
	ID        string    `json:"id"`
	BoardID   string    `json:"board_id"`
	Title     string    `json:"title"`
	Position  int       `json:"position"`
	Cards     []Card    `json:"cards"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type Card struct {
	ID          string `json:"id"`
	ColumnID    string `json:"column_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Order       int    `json:"order"`
}
