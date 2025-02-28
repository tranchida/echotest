package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"

	"github.com/tranchida/echotest/internal/handler"
	"github.com/tranchida/echotest/internal/middleware"
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
	mux.HandleFunc("GET /host.json", handler.HostInfoJsonHandler)

	return mux
}

// loggingMiddleware logs HTTP requests in a format similar to Echo's default logger


func main() {
	fmt.Println("open browser on : http://localhost:8080")

	server := &http.Server{
		Addr:    ":8080",
		Handler: middleware.LoggingMiddleware(setupServer()),
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
