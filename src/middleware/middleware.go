package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs information about incoming requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Process the request
		next.ServeHTTP(w, r)

		// Log the request details
		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			time.Since(start),
		)
	})
}
