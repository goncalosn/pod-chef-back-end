package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	handlers "pod-chef-back-end/handlers"
	kubernetesServices "pod-chef-back-end/internal/core/services/kubernetes"
	repositories "pod-chef-back-end/repositories/kubernetes"
)

func main() {

	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	kubernetesRepository := repositories.NewKubernetesRepository()
	kubernetesService := kubernetesServices.NewKubernetesService(kubernetesRepository)
	handler := handlers.NewHTTPHandler(kubernetesService)

	handlers.DeploymentsHandler(e, handler)
	handlers.NamespacesHandler(e, handler)
	handlers.NodesHandler(e, handler)

	e.GET("/api", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))

}
