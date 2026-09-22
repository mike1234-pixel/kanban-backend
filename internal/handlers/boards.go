package handlers

import (
	"encoding/json"
	"net/http"

	"kanban-backend/internal/models"
	"kanban-backend/internal/store"
)

type BoardHandler struct {
	Store store.Store
}

func (h *BoardHandler) CreateBoard(w http.ResponseWriter, r *http.Request) {
	var board models.Board
	if err := json.NewDecoder(r.Body).Decode(&board); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	if board.ID == "" || board.Title == "" {
		http.Error(w, "id and title are required", http.StatusBadRequest)
		return
	}

	if err := h.Store.SaveBoard(board); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(board)
}

func (h *BoardHandler) ListBoards(w http.ResponseWriter, r *http.Request) {
	boards, err := h.Store.GetBoards()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(boards)
}

func (h *BoardHandler) GetBoard(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	board, err := h.Store.GetBoardByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if board.ID == "" {
		http.Error(w, "Board not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(board)
}

func (h *BoardHandler) DeleteBoard(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	deleted, err := h.Store.DeleteBoard(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !deleted {
		http.Error(w, "Board not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
