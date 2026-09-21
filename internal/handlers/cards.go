package handlers

import (
	"encoding/json"
	"net/http"

	"kanban-backend/internal/models"
	"kanban-backend/internal/store"
)

// CardHandler holds dependencies needed by HTTP routes
type CardHandler struct {
	Store *store.MemoryStore
}

// CreateCard handles POST /cards
func (h *CardHandler) CreateCard(w http.ResponseWriter, r *http.Request) {
	var card models.Card

	// 1. Decode JSON from request body into the 'card' struct
	if err := json.NewDecoder(r.Body).Decode(&card); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	// 2. Save card to our memory store
	h.Store.SaveCard(card)

	// 3. Return JSON response with 201 Created status
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(card)
}

// ListCards handles GET /cards
func (h *CardHandler) ListCards(w http.ResponseWriter, r *http.Request) {
	// 1. Fetch cards from memory store
	cards := h.Store.GetCards()

	// 2. Return cards as JSON with 200 OK
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cards)
}

// GetCard handles GET /cards/{id}
func (h *CardHandler) GetCard(w http.ResponseWriter, r *http.Request) {
	// r.PathValue extracts wildcards matched in the route pattern (Go 1.22+)
	id := r.PathValue("id")

	card, found := h.Store.GetCardByID(id)
	if !found {
		http.Error(w, "Card not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(card)
}

// UpdateCard handles PUT /cards/{id}
func (h *CardHandler) UpdateCard(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var card models.Card
	if err := json.NewDecoder(r.Body).Decode(&card); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	success := h.Store.UpdateCard(id, card)
	if !success {
		http.Error(w, "Card not found", http.StatusNotFound)
		return
	}

	// Fetch updated card to return in response
	updatedCard, _ := h.Store.GetCardByID(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedCard)
}

// DeleteCard handles DELETE /cards/{id}
func (h *CardHandler) DeleteCard(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	success := h.Store.DeleteCard(id)
	if !success {
		http.Error(w, "Card not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 No Content
}
