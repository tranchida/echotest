package handler

import (
	"net/http"

	"github.com/tranchida/echotest/internal/component"
	"github.com/tranchida/echotest/internal/model"
)

// IndexHandler handles the / route.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	Render(w, r, http.StatusOK, component.Index())
}

// HostInfoHandler handles the /host route.
func HostInfoHandler(w http.ResponseWriter, r *http.Request) {
	hostInfo, err := model.GetHostInfo()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	Render(w, r, http.StatusOK, component.HostDisplay(hostInfo))
}
