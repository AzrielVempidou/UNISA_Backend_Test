package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"UNISA_Server/utils"
)

// Define custom types for context keys
type contextKey string

const (
	userIDKey contextKey = "userID"
	roleKey   contextKey = "role"
)

var jwtSecret = []byte("your_secret_key")

// AuthMiddleware verifies the JWT token and extracts user information
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("access_token")
		if authHeader == "" {
			utils.ErrorResponse(w, http.StatusUnauthorized, "Missing token")
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			utils.ErrorResponse(w, http.StatusUnauthorized, "Invalid token")
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			utils.ErrorResponse(w, http.StatusUnauthorized, "Invalid token claims")
			return
		}

		ctx := context.WithValue(r.Context(), userIDKey, claims["id"])
		ctx = context.WithValue(ctx, roleKey, claims["role"])
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// AdminMiddleware checks if the user has admin role
func AdminMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		role, ok := r.Context().Value(roleKey).(string)
		if !ok || role != "admin" {
			utils.ErrorResponse(w, http.StatusForbidden, "Access denied")
			return
		}
		next.ServeHTTP(w, r)
	}
}
