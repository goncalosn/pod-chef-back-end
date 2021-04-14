package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	handlers "pod-chef-back-end/handlers"
	services "pod-chef-back-end/internal/core/services"
	repositories "pod-chef-back-end/repositories/kubernetes"
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

	kubernetesRepository := repositories.KubernetesRepository()
	nodeServices := services.NodeServices(kubernetesRepository.Nodes)
	podServices := services.PodServices(kubernetesRepository.Pods)
	deploymentServices := services.DeploymentServices(kubernetesRepository.Deployments)
	handlers.NodeHandler(e, nodeServices)
	handlers.PodHandler(e, podServices)
	handlers.DeploymentHandler(e, deploymentServices)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))

}
