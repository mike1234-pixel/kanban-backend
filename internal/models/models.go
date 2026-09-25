package models

import "time"

type Board struct {
	ID        string    `json:"id" validate:"required,uuid"`
	Title     string    `json:"title" validate:"required,min=3,max=100"`
	Columns   []Column  `json:"columns" validate:"dive"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type Column struct {
	ID        string    `json:"id" validate:"required,uuid"`
	BoardID   string    `json:"board_id" validate:"required,uuid"`
	Title     string    `json:"title" validate:"required,min=1,max=100"`
	Position  int       `json:"position" validate:"gte=0"`
	Cards     []Card    `json:"cards,omitempty" validate:"dive"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type Card struct {
	ID          string `json:"id" validate:"required,uuid"`
	ColumnID    string `json:"column_id" validate:"required,uuid"`
	Title       string `json:"title" validate:"required,min=1,max=150"`
	Description string `json:"description" validate:"max=1000"`
	Order       int    `json:"order" validate:"gte=0"`
}
