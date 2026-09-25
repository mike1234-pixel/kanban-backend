package store

import (
	"sync"

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
	DeleteColumn(id string) (bool, error)

	// Board operations
	SaveBoard(board models.Board) error
	GetBoards() ([]models.Board, error)
	GetBoardByID(id string) (models.Board, error)
	DeleteBoard(id string) (bool, error)
}

type MemoryStore struct {
	mu      sync.RWMutex
	cards   map[string]models.Card
	columns map[string]models.Column
	boards  map[string]models.Board
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		cards:   make(map[string]models.Card),
		columns: make(map[string]models.Column),
		boards:  make(map[string]models.Board),
	}
}

// Card operations
func (s *MemoryStore) SaveCard(card models.Card) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.cards[card.ID] = card
	return nil
}

func (s *MemoryStore) GetCards() ([]models.Card, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	list := make([]models.Card, 0, len(s.cards))
	for _, card := range s.cards {
		list = append(list, card)
	}
	return list, nil
}

func (s *MemoryStore) GetCardByID(id string) (models.Card, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	card, exists := s.cards[id]
	if !exists {
		return models.Card{}, nil
	}
	return card, nil
}

func (s *MemoryStore) UpdateCard(id string, updated models.Card) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.cards[id]; !exists {
		return false, nil
	}

	updated.ID = id
	s.cards[id] = updated
	return true, nil
}

func (s *MemoryStore) DeleteCard(id string) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.cards[id]; !exists {
		return false, nil
	}

	delete(s.cards, id)
	return true, nil
}

// Column operations
func (s *MemoryStore) SaveColumn(col models.Column) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.columns[col.ID] = col
	return nil
}

func (s *MemoryStore) GetColumnsWithCards(boardID string) ([]models.Column, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	cols := make([]models.Column, 0)
	for _, col := range s.columns {
		if col.BoardID == boardID {
			// Populate nested cards for each column
			colCards := make([]models.Card, 0)
			for _, card := range s.cards {
				if card.ColumnID == col.ID {
					colCards = append(colCards, card)
				}
			}
			col.Cards = colCards
			cols = append(cols, col)
		}
	}
	return cols, nil
}

func (s *MemoryStore) GetColumnByID(id string) (models.Column, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	col, exists := s.columns[id]
	if !exists {
		return models.Column{}, nil
	}
	return col, nil
}

func (s *MemoryStore) DeleteColumn(id string) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.columns[id]; !exists {
		return false, nil
	}

	delete(s.columns, id)
	return true, nil
}

// Board operations
func (s *MemoryStore) SaveBoard(board models.Board) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.boards[board.ID] = board
	return nil
}

func (s *MemoryStore) GetBoards() ([]models.Board, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	list := make([]models.Board, 0, len(s.boards))
	for _, board := range s.boards {
		list = append(list, board)
	}
	return list, nil
}

func (s *MemoryStore) GetBoardByID(id string) (models.Board, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	board, exists := s.boards[id]
	if !exists {
		return models.Board{}, nil
	}
	return board, nil
}

func (s *MemoryStore) DeleteBoard(id string) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.boards[id]; !exists {
		return false, nil
	}

	delete(s.boards, id)
	return true, nil
}
