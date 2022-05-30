package middlewares

import "net/http"

// Adds a Content-Type applicatoin/json to the request header
func AddJsonContentType(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}
