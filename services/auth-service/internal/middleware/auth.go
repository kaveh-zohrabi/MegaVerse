package middleware

import (
	"context"
	"net/http"
	"strings"

	"megaverse.dev/services/auth-service/internal/service"
)

type contextKey string

const (
	UserIDKey contextKey = "user_id"
	EmailKey  contextKey = "email"
	RoleKey   contextKey = "role"
)

func AuthMiddleware(authService *service.AuthService) func(http.Handler) http.Handler {
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

			claims, err := authService.ValidateToken(parts[1])
			if err != nil {
				http.Error(w, `{"error":"Invalid token"}`, http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
			ctx = context.WithValue(ctx, EmailKey, claims.Email)
			ctx = context.WithValue(ctx, RoleKey, claims.Role)

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
