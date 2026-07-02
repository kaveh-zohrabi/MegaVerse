package middleware

import (
	"context"
	"net/http"
	"strings"
)

type contextKey string

const (
	UserIDKey contextKey = "user_id"
	EmailKey  contextKey = "email"
	RoleKey   contextKey = "role"
)

func AuthMiddleware(authServiceURL string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, `{"error":"Authorization header required"}`, http.StatusUnauthorized)
				return
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, `{"error":"Invalid authorization format"}`, http.StatusUnauthorized)
				return
			}

			// In production, call auth service to validate
			// For now, extract user ID from token claims
			ctx := context.WithValue(r.Context(), UserIDKey, r.Header.Get("X-User-ID"))
			ctx = context.WithValue(ctx, EmailKey, r.Header.Get("X-User-Email"))
			ctx = context.WithValue(ctx, RoleKey, r.Header.Get("X-User-Role"))

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUserID(r *http.Request) string {
	if val := r.Context().Value(UserIDKey); val != nil {
		return val.(string)
	}
	return ""
}

func GetEmail(r *http.Request) string {
	if val := r.Context().Value(EmailKey); val != nil {
		return val.(string)
	}
	return ""
}

func GetRole(r *http.Request) string {
	if val := r.Context().Value(RoleKey); val != nil {
		return val.(string)
	}
	return ""
}
