package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tranchida/echotest/internal/handlers"
)

func newEcho() (*echo.Echo, error) {

	e := echo.New()
	e.HideBanner = false

	e.Use(middleware.Gzip())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "${remote_ip} - - [${time_custom}] \"${method} ${uri} ${protocol}\" ${status} ${bytes_in} ${bytes_out} ${latency_human}\n",
		CustomTimeFormat: "02/Jan/2006:15:04:05 -0700",
	}))

	e.GET("/", handlers.IndexHandler)

	e.GET("/host", handlers.HostHandler)

	return e, nil
}

func main() {

	e, err := newEcho()
	if err != nil {
		panic(err)
	}

	if err := e.Start(":8080"); err != nil {
		panic(err)
	}

}
