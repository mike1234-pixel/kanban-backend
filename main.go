package main

import (
	"fmt"
	"net/http"

	"kanban-backend/internal/handlers"
	"kanban-backend/internal/store"
)

func main() {
	// 1. Initialize our in-memory data store
	memStore := store.NewMemoryStore()

	// 2. Pass the store to our card handler dependency
	cardHandler := &handlers.CardHandler{
		Store: memStore,
	}

	// 3. Create a new HTTP request multiplexer (router)
	mux := http.NewServeMux()

	// 4. Register our route handlers
	mux.HandleFunc("POST /cards", cardHandler.CreateCard)
	mux.HandleFunc("GET /cards", cardHandler.ListCards)
	mux.HandleFunc("GET /cards/{id}", cardHandler.GetCard)
	mux.HandleFunc("PUT /cards/{id}", cardHandler.UpdateCard)
	mux.HandleFunc("DELETE /cards/{id}", cardHandler.DeleteCard)

	// 5. Start the HTTP server on port 8080
	fmt.Println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
