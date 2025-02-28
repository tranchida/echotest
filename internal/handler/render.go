package handler

import (
	"net/http"

	"github.com/a-h/templ"
)

func Render(w http.ResponseWriter, r *http.Request, statusCode int, t templ.Component) {
	w.WriteHeader(statusCode)
	t.Render(r.Context(), w)
}