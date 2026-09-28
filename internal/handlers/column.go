package handlers

import (
	"encoding/json"
	"net/http"

	"kanban-backend/internal/models"
	"kanban-backend/internal/store"

	"github.com/go-playground/validator/v10"
)

type ColumnHandler struct {
	Store    store.Store
	Validate *validator.Validate
}

// #region Create Column
func (h *ColumnHandler) CreateColumn(w http.ResponseWriter, r *http.Request) {
	var col models.Column
	if err := json.NewDecoder(r.Body).Decode(&col); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	if err := h.Validate.Struct(col); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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

//#endregion

// #region List Columns
func (h *ColumnHandler) ListColumns(w http.ResponseWriter, r *http.Request) {
	boardID := r.URL.Query().Get("board_id")
	if boardID == "" {
		http.Error(w, "board_id query parameter is required", http.StatusBadRequest)
		return
	}

	columns, err := h.Store.GetColumnsWithCards(boardID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(columns)
}

//#endregion

// #region Delete Column
func (h *ColumnHandler) DeleteColumn(w http.ResponseWriter, r *http.Request) {
	// Extract 'id' from path parameter (Go 1.22+ net/http pattern routing)
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "Column ID is required", http.StatusBadRequest)
		return
	}

	deleted, err := h.Store.DeleteColumn(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if !deleted {
		http.Error(w, "Column not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

//#endregion
