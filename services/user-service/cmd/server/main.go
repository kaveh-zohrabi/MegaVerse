package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"

	"megaverse.dev/services/user-service/internal/handler"
	"megaverse.dev/services/user-service/internal/middleware"
	"megaverse.dev/services/user-service/internal/repository"
	"megaverse.dev/services/user-service/internal/service"
)

func main() {
	dbHost := getEnv("DATABASE_HOST", "localhost")
	dbPort := getEnv("DATABASE_PORT", "3306")
	dbName := getEnv("DATABASE_NAME", "megaverse")
	dbUser := getEnv("DATABASE_USER", "megaverse")
	dbPass := getEnv("DATABASE_PASSWORD", "password")
	authServiceURL := getEnv("AUTH_SERVICE_URL", "http://localhost:8081")
	port := getEnv("PORT", "8082")

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

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	router := mux.NewRouter()

	// Health check
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"healthy","service":"user-service"}`))
	}).Methods("GET")

	// Public routes
	router.HandleFunc("/users/{id}", userHandler.GetProfile).Methods("GET")
	router.HandleFunc("/users/{id}/followers", userHandler.GetFollowers).Methods("GET")
	router.HandleFunc("/users/{id}/following", userHandler.GetFollowing).Methods("GET")

	// Protected routes
	protected := router.PathPrefix("/users").Subrouter()
	protected.Use(middleware.AuthMiddleware(authServiceURL))
	protected.HandleFunc("/me", userHandler.UpdateProfile).Methods("PUT")
	protected.HandleFunc("/{id}/follow", userHandler.Follow).Methods("POST")
	protected.HandleFunc("/{id}/unfollow", userHandler.Unfollow).Methods("POST")

	addr := fmt.Sprintf(":%s", port)
	log.Printf("User service starting on %s", addr)
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
