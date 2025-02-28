package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggerMiddleware logs HTTP requests in a format similar to Echo's logger
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Create a wrapper for the response writer to capture the status code
		wrapper := &responseWriterWrapper{ResponseWriter: w, statusCode: http.StatusOK}

		// Process the request
		next.ServeHTTP(wrapper, r)

		// Log the request details
		log.Printf("%s - - [%s] \"%s %s %s\" %d %s\n",
			r.RemoteAddr,
			time.Now().Format("02/Jan/2006:15:04:05 -0700"),
			r.Method,
			r.URL.Path,
			r.Proto,
			wrapper.statusCode,
			time.Since(start),
		)
	})
}

// responseWriterWrapper captures the status code of the response
type responseWriterWrapper struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader captures the status code
func (w *responseWriterWrapper) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
