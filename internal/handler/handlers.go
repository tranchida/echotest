package handler

import (
	"html/template"
	"net/http"

	"github.com/tranchida/echotest/internal/model" // Updated import path
)

// IndexHandler handles the / route.
func IndexHandler(template *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if template.ExecuteTemplate(w, "index.html", nil) != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

// HostInfoHandler handles the /host route.
func HostInfoHandler(template *template.Template) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		hostInfo, err := model.GetHostInfo()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if template.ExecuteTemplate(w, "host.html", hostInfo) != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}
