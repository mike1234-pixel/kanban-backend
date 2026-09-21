package handlers

import (
	"encoding/json"
	"net/http"

	"kanban-backend/internal/models"
	"kanban-backend/internal/store"
)

// CardHandler holds dependencies needed by HTTP routes
type CardHandler struct {
	Store store.Store
}

// CreateCard handles POST /cards
func (h *CardHandler) CreateCard(w http.ResponseWriter, r *http.Request) {
	var card models.Card

	// 1. Decode JSON from request body into the 'card' struct
	if err := json.NewDecoder(r.Body).Decode(&card); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	// 2. Save card to store and check for errors
	if err := h.Store.SaveCard(card); err != nil {
		http.Error(w, "Failed to save card", http.StatusInternalServerError)
		return
	}

	// 3. Return JSON response with 201 Created status
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(card)
}

// ListCards handles GET /cards
func (h *CardHandler) ListCards(w http.ResponseWriter, r *http.Request) {
	// 1. Fetch cards from store
	cards, err := h.Store.GetCards()
	if err != nil {
		http.Error(w, "Failed to fetch cards", http.StatusInternalServerError)
		return
	}

	// 2. Return cards as JSON with 200 OK
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cards)
}

// GetCard handles GET /cards/{id}
func (h *CardHandler) GetCard(w http.ResponseWriter, r *http.Request) {
	// r.PathValue extracts wildcards matched in the route pattern (Go 1.22+)
	id := r.PathValue("id")

	card, err := h.Store.GetCardByID(id)
	if err != nil {
		http.Error(w, "Failed to fetch card", http.StatusInternalServerError)
		return
	}
	if card.ID == "" {
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

	// 1. Update the card in the store
	success, err := h.Store.UpdateCard(id, card)
	if err != nil {
		http.Error(w, "Failed to update card", http.StatusInternalServerError)
		return
	}
	if !success {
		http.Error(w, "Card not found", http.StatusNotFound)
		return
	}

	// 2. Fetch updated card to return in response
	updatedCard, err := h.Store.GetCardByID(id)
	if err != nil {
		http.Error(w, "Failed to fetch updated card", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedCard)
}

// DeleteCard handles DELETE /cards/{id}
func (h *CardHandler) DeleteCard(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	success, err := h.Store.DeleteCard(id)
	if err != nil {
		http.Error(w, "Failed to delete card", http.StatusInternalServerError)
		return
	}
	if !success {
		http.Error(w, "Card not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 No Content
}
