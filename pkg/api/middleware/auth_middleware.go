package middleware

import "net/http"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Implement your authentication logic here
		// If unauthorized, you can return a 401 status code:
		// http.Error(w, "Unauthorized", http.StatusUnauthorized)
		// return

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
