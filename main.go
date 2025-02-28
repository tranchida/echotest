package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/tranchida/echotest/internal/handler"
)

//go:embed static
var contentFS embed.FS

func setupServer() *http.ServeMux {
	mux := http.NewServeMux()

	// Serve static files
	fileServer := http.FileServer(http.FS(contentFS))
	mux.Handle("GET /static/", fileServer)

	// Routes
	mux.HandleFunc("GET /", handler.IndexHandler)
	mux.HandleFunc("GET /host", handler.HostInfoHandler)

	return mux
}

// loggingMiddleware logs HTTP requests in a format similar to Echo's default logger
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Wrap ResponseWriter to capture status code
		rw := &responseWriter{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(rw, r)

		log.Printf("%s - - [%s] \"%s %s %s\" %d %v\n",
			r.RemoteAddr,
			start.Format("02/Jan/2006:15:04:05 -0700"),
			r.Method,
			r.URL.Path,
			r.Proto,
			rw.status,
			time.Since(start),
		)
	})
}

// responseWriter wraps http.ResponseWriter to capture status code
type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func main() {
	fmt.Println("open browser on : http://localhost:8080")

	server := &http.Server{
		Addr:    ":8080",
		Handler: loggingMiddleware(setupServer()),
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
