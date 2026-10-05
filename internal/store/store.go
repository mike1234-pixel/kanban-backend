package store

import (
	"kanban-backend/internal/models"
)

type Store interface {
	// Card operations
	SaveCard(card models.Card) error
	GetCards() ([]models.Card, error)
	GetCardByID(id string) (models.Card, error)
	UpdateCard(id string, updated models.Card) (bool, error)
	DeleteCard(id string) (bool, error)

	// Column operations
	SaveColumn(col models.Column) error
	GetColumnsWithCards(boardID string) ([]models.Column, error)
	GetColumnByID(id string) (models.Column, error)
	UpdateColumn(col models.Column) (bool, error)
	DeleteColumn(id string) (bool, error)

	// Board operations
	SaveBoard(board models.Board) error
	GetBoards() ([]models.Board, error)
	GetBoardByID(id string) (models.Board, error)
	UpdateBoard(board models.Board) (bool, error)
	DeleteBoard(id string) (bool, error)
}
