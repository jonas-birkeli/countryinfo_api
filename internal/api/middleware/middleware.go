package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Middleware type definition
type Middleware func(http.Handler) http.Handler

// Chain applies middlewares in order
func Chain(h http.Handler, middlewares ...Middleware) http.Handler {
	for _, m := range middlewares {
		h = m(h)
	}
	return h
}

// Logger logs request details to standard out
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Log the request
		log.Printf(
			"Started %s %s",
			r.Method,
			r.URL.Path,
		)

		next.ServeHTTP(w, r)

		// Log the completion
		log.Printf(
			"Completed in %v",
			time.Since(start),
		)
	})
}

// CORS adds headers for CORS service
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Allow explicit preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// RequestID adds a request ID to each request
func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Could use UUID but staying with a standard library
		requestID := time.Now().UnixNano()
		w.Header().Set("X-Request-ID", strconv.FormatInt(requestID, 10))

		next.ServeHTTP(w, r)
	})
}

// Recover middleware to handle panics
func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("panic: %v", err)
				log.Printf("panic: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
