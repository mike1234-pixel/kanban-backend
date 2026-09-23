package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"kanban-backend/internal/handlers"
	"kanban-backend/internal/store"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, relying on environment variables")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser, dbPassword, dbHost, dbPort, dbName, dbSSLMode)

	dbStore, err := store.NewPostgresStore(connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	cardHandler := &handlers.CardHandler{
		Store: dbStore,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /cards", cardHandler.CreateCard)
	mux.HandleFunc("GET /cards", cardHandler.ListCards)
	mux.HandleFunc("GET /cards/{id}", cardHandler.GetCard)
	mux.HandleFunc("PUT /cards/{id}", cardHandler.UpdateCard)
	mux.HandleFunc("DELETE /cards/{id}", cardHandler.DeleteCard)

	columnHandler := &handlers.ColumnHandler{Store: dbStore}

	mux.HandleFunc("POST /columns", columnHandler.CreateColumn)
	mux.HandleFunc("GET /columns", columnHandler.ListColumns)

	boardHandler := &handlers.BoardHandler{Store: dbStore}

	mux.HandleFunc("POST /boards", boardHandler.CreateBoard)
	mux.HandleFunc("GET /boards", boardHandler.ListBoards)
	mux.HandleFunc("GET /boards/{id}", boardHandler.GetBoard)
	mux.HandleFunc("DELETE /boards/{id}", boardHandler.DeleteBoard)

	fmt.Println("Server running on http://localhost:8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
