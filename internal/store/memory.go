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
	GetColumnsWithCards(boardID string) ([]models.Column, error) // <-- Update signature here
	GetColumnByID(id string) (models.Column, error)
	DeleteColumn(id string) (bool, error)

	// Board operations
	SaveBoard(board models.Board) error
	GetBoards() ([]models.Board, error)
	GetBoardByID(id string) (models.Board, error)
	DeleteBoard(id string) (bool, error)
}

// this whole block would be a class in TS
type MemoryStore struct {
	mu    sync.RWMutex // "Read/Write Mutex". It acts like a digital traffic light controlling access to the cards map so two requests don't modify it simultaneously.
	cards map[string]models.Card
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		cards: make(map[string]models.Card),
	}
}

// SaveCard adds or overwrites a card in the store
func (s *MemoryStore) SaveCard(card models.Card) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.cards[card.ID] = card
	return nil
}

// GetCards retrieves all cards from the store
func (s *MemoryStore) GetCards() ([]models.Card, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	list := make([]models.Card, 0, len(s.cards))
	for _, card := range s.cards {
		list = append(list, card)
	}
	return list, nil
}

// GetCardByID retrieves a single card by its ID
func (s *MemoryStore) GetCardByID(id string) (models.Card, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	card, exists := s.cards[id]
	if !exists {
		return models.Card{}, nil
	}
	return card, nil
}

// UpdateCard updates an existing card in the store
func (s *MemoryStore) UpdateCard(id string, updated models.Card) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.cards[id]; !exists {
		return false, nil
	}

	// Ensure the ID in the card struct matches the URL ID
	updated.ID = id
	s.cards[id] = updated
	return true, nil
}

// DeleteCard removes a card from the store by its ID
func (s *MemoryStore) DeleteCard(id string) (bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.cards[id]; !exists {
		return false, nil
	}

	delete(s.cards, id)
	return true, nil
}
