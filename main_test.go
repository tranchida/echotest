package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestHostEndpoint(t *testing.T) {
	handler := loggingMiddleware(setupServer())

	req := httptest.NewRequest(http.MethodGet, "/host", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status %d but got %d", http.StatusOK, rec.Code)
	}

	hostname, err := os.Hostname()
	if err != nil {
		t.Fatalf("failed to get hostname: %v", err)
	}

	body := rec.Body.String()
	if !strings.Contains(body, hostname) {
		t.Errorf("response body does not contain hostname, got: %s", body)
	}
}

func TestStaticFileNotFound(t *testing.T) {
	handler := loggingMiddleware(setupServer())

	req := httptest.NewRequest(http.MethodGet, "/nonexistentfile.txt", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("Expected status %d but got %d", http.StatusNotFound, rec.Code)
	}
}
