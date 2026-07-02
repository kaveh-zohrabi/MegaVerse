package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"

	"megaverse.dev/services/api-gateway/internal/config"
	"megaverse.dev/services/api-gateway/internal/handler"
	"megaverse.dev/services/api-gateway/internal/middleware"
)

func main() {
	cfg := config.Load()

	// Define service URLs
	services := map[string]*url.URL{
		"auth":   {Scheme: "http", Host: "localhost:8081"},
		"user":   {Scheme: "http", Host: "localhost:8082"},
		"social": {Scheme: "http", Host: "localhost:8083"},
	}

	// Override from env
	if authURL := cfg.AuthServiceURL; authURL != "" {
		if u, err := url.Parse(authURL); err == nil {
			services["auth"] = u
		}
	}

	proxyHandler := handler.NewProxyHandler(services)
	rateLimiter := middleware.NewRateLimiter(cfg.RateLimit, cfg.RateWindowSec)

	router := mux.NewRouter()

	// Middleware
	router.Use(middleware.CORSMiddleware)
	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.RateLimitMiddleware(rateLimiter))

	// Health check
	router.HandleFunc("/health", proxyHandler.HealthCheck).Methods("GET")

	// Auth routes (public)
	router.PathPrefix("/auth").HandlerFunc(proxyHandler.Proxy("auth"))

	// User routes
	router.PathPrefix("/users").HandlerFunc(proxyHandler.Proxy("user"))

	// Social routes
	router.PathPrefix("/posts").HandlerFunc(proxyHandler.Proxy("social"))
	router.PathPrefix("/feed").HandlerFunc(proxyHandler.Proxy("social"))

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("API Gateway starting on %s", addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
