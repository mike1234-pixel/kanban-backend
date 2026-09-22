package store

import (
	"database/sql"
	"fmt"

	"kanban-backend/internal/models"

	_ "github.com/lib/pq" // Driver side-effect import to register PostgreSQL with database/sql
)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore(connStr string) (*PostgresStore, error) {
	// 1. sql.Open initializes the connection pool manager (no TCP handshake yet)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// 2. Ping forces an actual TCP connection to PostgreSQL on port 5432
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping postgres: %w", err)
	}

	return &PostgresStore{db: db}, nil
}

// SaveCard inserts a new card into the database
func (s *PostgresStore) SaveCard(card models.Card) error {
	query := `
		INSERT INTO cards (id, column_id, title, description, "order")
		VALUES ($1, $2, $3, $4, $5)
	`
	_, err := s.db.Exec(query, card.ID, card.ColumnID, card.Title, card.Description, card.Order)
	return err
}

// GetCards fetches all cards from the database
// GetCards fetches all cards from the database
func (s *PostgresStore) GetCards() ([]models.Card, error) {
	query := `SELECT id, column_id, title, description, "order" FROM cards`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close() // ALWAYS close rows to release the network connection back to the pool!

	var cards []models.Card
	for rows.Next() {
		var c models.Card
		if err := rows.Scan(&c.ID, &c.ColumnID, &c.Title, &c.Description, &c.Order); err != nil {
			return nil, err
		}
		cards = append(cards, c)
	}

	// CHECK FOR STREAMING/NETWORK ERRORS THAT OCCURRED DURING ITERATION
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}

	return cards, nil
}

// GetCardByID fetches a single card by its ID
func (s *PostgresStore) GetCardByID(id string) (models.Card, error) {
	query := `SELECT id, column_id, title, description, "order" FROM cards WHERE id = $1`

	var c models.Card
	err := s.db.QueryRow(query, id).Scan(&c.ID, &c.ColumnID, &c.Title, &c.Description, &c.Order)
	if err == sql.ErrNoRows {
		return models.Card{}, nil // Return empty card to signal not found
	}
	if err != nil {
		return models.Card{}, err
	}

	return c, nil
}

// UpdateCard updates an existing card's details
func (s *PostgresStore) UpdateCard(id string, updated models.Card) (bool, error) {
	query := `
		UPDATE cards 
		SET column_id = $1, title = $2, description = $3, "order" = $4
		WHERE id = $5
	`
	result, err := s.db.Exec(query, updated.ColumnID, updated.Title, updated.Description, updated.Order, id)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}

// DeleteCard removes a card by its ID
func (s *PostgresStore) DeleteCard(id string) (bool, error) {
	query := `DELETE FROM cards WHERE id = $1`
	result, err := s.db.Exec(query, id)
	if err != nil {
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rowsAffected > 0, nil
}

func (s *PostgresStore) SaveColumn(col models.Column) error {
	query := `INSERT INTO columns (id, title, position) VALUES ($1, $2, $3)`
	_, err := s.db.Exec(query, col.ID, col.Title, col.Position)
	return err
}

// GetColumnsWithCards performs a two-step load to construct the full board hierarchy
func (s *PostgresStore) GetColumnsWithCards() ([]models.Column, error) {
	// 1. Fetch all columns
	colQuery := `SELECT id, title, position FROM columns ORDER BY position ASC`
	colRows, err := s.db.Query(colQuery)
	if err != nil {
		return nil, err
	}
	defer colRows.Close()

	var columns []models.Column
	colMap := make(map[string]int) // Maps column ID to index in slice

	for colRows.Next() {
		var c models.Column
		c.Cards = []models.Card{} // Guarantee non-nil JSON array ([])
		if err := colRows.Scan(&c.ID, &c.Title, &c.Position); err != nil {
			return nil, err
		}
		colMap[c.ID] = len(columns)
		columns = append(columns, c)
	}
	if err := colRows.Err(); err != nil {
		return nil, err
	}

	// 2. Fetch all cards and distribute them to their respective columns
	cardQuery := `SELECT id, column_id, title, description, "order" FROM cards ORDER BY "order" ASC`
	cardRows, err := s.db.Query(cardQuery)
	if err != nil {
		return nil, err
	}
	defer cardRows.Close()

	for cardRows.Next() {
		var card models.Card
		if err := cardRows.Scan(&card.ID, &card.ColumnID, &card.Title, &card.Description, &card.Order); err != nil {
			return nil, err
		}
		if idx, exists := colMap[card.ColumnID]; exists {
			columns[idx].Cards = append(columns[idx].Cards, card)
		}
	}
	if err := cardRows.Err(); err != nil {
		return nil, err
	}

	return columns, nil
}

func (s *PostgresStore) GetColumnByID(id string) (models.Column, error) {
	query := `SELECT id, title, position FROM columns WHERE id = $1`
	var c models.Column
	err := s.db.QueryRow(query, id).Scan(&c.ID, &c.Title, &c.Position)
	if err == sql.ErrNoRows {
		return models.Column{}, nil
	}
	if err != nil {
		return models.Column{}, err
	}
	return c, nil
}

func (s *PostgresStore) DeleteColumn(id string) (bool, error) {
	query := `DELETE FROM columns WHERE id = $1`
	res, err := s.db.Exec(query, id)
	if err != nil {
		return false, err
	}
	rows, err := res.RowsAffected()
	return rows > 0, err
}
