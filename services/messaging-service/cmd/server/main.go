package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"

	"megaverse.dev/services/messaging-service/internal/handler"
	"megaverse.dev/services/messaging-service/internal/repository"
	"megaverse.dev/services/messaging-service/internal/service"
	ws "megaverse.dev/services/messaging-service/internal/websocket"
)

func main() {
	dbHost := getEnv("DATABASE_HOST", "localhost")
	dbPort := getEnv("DATABASE_PORT", "3306")
	dbName := getEnv("DATABASE_NAME", "megaverse")
	dbUser := getEnv("DATABASE_USER", "megaverse")
	dbPass := getEnv("DATABASE_PASSWORD", "password")
	port := getEnv("PORT", "8084")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	log.Println("Connected to MySQL")

	// Initialize
	repo := repository.NewMessagingRepository(db)
	svc := service.NewMessagingService(repo)
	h := handler.NewMessagingHandler(svc)
	hub := ws.NewHub()
	go hub.Run()

	router := mux.NewRouter()

	// Health check
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"healthy","service":"messaging-service"}`))
	}).Methods("GET")

	// WebSocket
	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws.HandleWebSocket(hub, w, r)
	})

	// REST API
	router.HandleFunc("/conversations", h.CreateConversation).Methods("POST")
	router.HandleFunc("/conversations", h.GetConversations).Methods("GET")
	router.HandleFunc("/conversations/{id}/messages", h.SendMessage).Methods("POST")
	router.HandleFunc("/conversations/{id}/messages", h.GetMessages).Methods("GET")

	addr := fmt.Sprintf(":%s", port)
	log.Printf("Messaging service starting on %s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
