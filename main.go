package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/tranchida/echotest/internal/handler"
	"github.com/tranchida/echotest/internal/middleware"
)

//go:embed static
var contentFS embed.FS

func setupServer() http.Handler {
	mux := http.NewServeMux()

	// Set up middleware
	loggedMux := middleware.LoggerMiddleware(mux)

	// Register handlers
	mux.HandleFunc("/", handler.IndexHandler)
	mux.HandleFunc("/host", handler.HostInfoHandler)

	// Serve static files
	fileServer := http.FileServer(http.FS(contentFS))
	mux.Handle("/static/", fileServer)

	return loggedMux
}

func main() {
	fmt.Println("Open browser on : http://localhost:8080")

	server := &http.Server{
		Addr:         ":8080",
		Handler:      setupServer(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
