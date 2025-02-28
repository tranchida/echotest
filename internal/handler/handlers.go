package handler

import (
	"net/http"

	"github.com/tranchida/echotest/internal/component"
	"github.com/tranchida/echotest/internal/model"
)

// IndexHandler handles the / route.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	err := Render(w, r, http.StatusOK, component.Index())
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

// HostInfoHandler handles the /host route.
func HostInfoHandler(w http.ResponseWriter, r *http.Request) {
	hostInfo, err := model.GetHostInfo()
	if err != nil {
		http.Error(w, "Error fetching host info: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = Render(w, r, http.StatusOK, component.HostDisplay(hostInfo))
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
