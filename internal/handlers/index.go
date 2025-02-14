package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tranchida/echotest/internal/components"
)

func IndexHandler(c echo.Context) error {
	return Render(c, http.StatusOK, components.Index())
}
