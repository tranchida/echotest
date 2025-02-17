package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tranchida/echotest/internal/model" // Updated import path
)

// HostInfoHandler handles the /host route.
func HostInfoHandler(c echo.Context) error {
	hostInfo, err := model.GetHostInfo()
	if err != nil {
		return err // Proper error handling; return the error to Echo
	}
	return c.Render(http.StatusOK, "main", hostInfo)
}

