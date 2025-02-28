package handler

import (
	"encoding/json"
	"net/http"

	"github.com/a-h/templ"
)

func Render(w http.ResponseWriter, r *http.Request, statusCode int, t templ.Component) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(statusCode)
	t.Render(r.Context(), w)
}


func RenderJson(w http.ResponseWriter, r *http.Request, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}