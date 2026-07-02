package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"

	"megaverse.dev/services/auth-service/internal/handler"
	"megaverse.dev/services/auth-service/internal/middleware"
	"megaverse.dev/services/auth-service/internal/repository"
	"megaverse.dev/services/auth-service/internal/service"
)

func main() {
	// Load configuration from environment
	dbHost := getEnv("DATABASE_HOST", "localhost")
	dbPort := getEnv("DATABASE_PORT", "3306")
	dbName := getEnv("DATABASE_NAME", "megaverse")
	dbUser := getEnv("DATABASE_USER", "megaverse")
	dbPass := getEnv("DATABASE_PASSWORD", "password")
	jwtSecret := getEnv("JWT_SECRET", "dev-secret-change-in-production")
	port := getEnv("PORT", "8081")

	// Connect to MySQL
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

	// Initialize layers
	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo, jwtSecret, 24*time.Hour, 7*24*time.Hour)
	authHandler := handler.NewAuthHandler(authService)

	// Setup router
	router := mux.NewRouter()

	// Public routes
	router.HandleFunc("/auth/register", authHandler.Register).Methods("POST")
	router.HandleFunc("/auth/login", authHandler.Login).Methods("POST")

	// Protected routes
	protected := router.PathPrefix("/auth").Subrouter()
	protected.Use(middleware.AuthMiddleware(authService))
	protected.HandleFunc("/validate", authHandler.Validate).Methods("GET")

	// Health check
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"healthy","service":"auth-service"}`))
	}).Methods("GET")

	// Start server
	addr := fmt.Sprintf(":%s", port)
	log.Printf("Auth service starting on %s", addr)
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
