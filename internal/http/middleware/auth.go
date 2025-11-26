package middleware

import (
	"net/http"
	"strings"
)

// Auth middleware skeleton: valida header Authorization.
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" { http.Error(w, "missing auth", http.StatusUnauthorized); return }
		if !strings.HasPrefix(auth, "Bearer ") { http.Error(w, "invalid auth", http.StatusUnauthorized); return }
		// parse token omitted
		next.ServeHTTP(w, r)
	})
}
