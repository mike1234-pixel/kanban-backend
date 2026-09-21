package models

type Board struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Column struct {
	ID      string `json:"id"`
	BoardID string `json:"board_id"`
	Title   string `json:"title"`
	Order   int    `json:"order"`
}

type Card struct {
	ID          string `json:"id"`
	ColumnID    string `json:"column_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Order       int    `json:"order"`
}
