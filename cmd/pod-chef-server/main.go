package main

import (
	http "pod-chef-back-end/pkg/http"

	kube "pod-chef-back-end/pkg/kubernetes"

	. "pod-chef-back-end/pkg/builders"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:1323/"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	kubernetesClient := kube.Client()

	services := BuildServices(kubernetesClient)
	interactors := BuildInteractors(services)
	http.BuildHandlers(e, interactors)

	e.Logger.Fatal(e.Start(":1323"))

}
