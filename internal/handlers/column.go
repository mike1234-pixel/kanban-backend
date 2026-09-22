package handlers

import (
	"encoding/json"
	"net/http"

	"kanban-backend/internal/models"
	"kanban-backend/internal/store"
)

type ColumnHandler struct {
	Store store.Store
}

func (h *ColumnHandler) CreateColumn(w http.ResponseWriter, r *http.Request) {
	var col models.Column
	if err := json.NewDecoder(r.Body).Decode(&col); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	if col.ID == "" || col.Title == "" {
		http.Error(w, "id and title are required", http.StatusBadRequest)
		return
	}

	if err := h.Store.SaveColumn(col); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(col)
}

func (h *ColumnHandler) ListColumns(w http.ResponseWriter, r *http.Request) {
	boardID := r.URL.Query().Get("board_id")
	if boardID == "" {
		boardID = "board-1" // Default board fallback
	}

	columns, err := h.Store.GetColumnsWithCards(boardID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(columns)
}
