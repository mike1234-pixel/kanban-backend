package store

import (
	"kanban-backend/internal/models"
	"sync"
)

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

// public methods. Capitalising method names makes them public.
func (s *MemoryStore) SaveCard(card models.Card) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.cards[card.ID] = card
}

func (s *MemoryStore) GetCards() []models.Card {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]models.Card, 0, len(s.cards))
	for _, card := range s.cards {
		list = append(list, card)
	}
	return list
}

// GetCardByID retrieves a single card by its ID
func (s *MemoryStore) GetCardByID(id string) (models.Card, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	card, exists := s.cards[id]
	return card, exists
}

// UpdateCard updates an existing card in the store
func (s *MemoryStore) UpdateCard(id string, updated models.Card) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.cards[id]; !exists {
		return false
	}

	// Ensure the ID in the card struct matches the URL ID
	updated.ID = id
	s.cards[id] = updated
	return true
}

// DeleteCard removes a card from the store by its ID
func (s *MemoryStore) DeleteCard(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.cards[id]; !exists {
		return false
	}

	delete(s.cards, id)
	return true
}
