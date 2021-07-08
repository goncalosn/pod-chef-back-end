package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/app", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/app/headers", headers)
	e.Logger.Fatal(e.Start(":8080"))
}

func headers(c echo.Context) error {
	return c.JSONPretty(http.StatusOK, c.Request().Header, " ")

}
