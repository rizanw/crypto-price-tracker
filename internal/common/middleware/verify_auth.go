package middleware

import (
	"context"
	"crypto-tracker/internal/common/constants"
	"crypto-tracker/internal/common/session"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
)

func VerifyAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// get auth token from header
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Missing authorization header", http.StatusUnauthorized)
			return
		}

		tokenString = tokenString[len("Bearer "):]

		// verify jwt token
		err := verifyToken(tokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// verify session expiry
		s, err := verifySession(tokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// pass session to context
		ctx := r.Context()
		ctx = context.WithValue(ctx, "session", s)
		ctx = context.WithValue(ctx, "token", tokenString)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return constants.JWTSecretKey, nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

func verifySession(tokenString string) (session.Session, error) {
	userSession, exists := session.Sessions[tokenString]
	if !exists {
		// if session token is not present in session map, return an unauthorized
		return userSession, fmt.Errorf("invalid token")
	}

	// if session is expired delete session map, return expired
	if userSession.IsExpired() {
		delete(session.Sessions, tokenString)
		return userSession, fmt.Errorf("session is expired")
	}

	return userSession, nil
}
