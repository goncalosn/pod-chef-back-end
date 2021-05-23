package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	handlers "pod-chef-back-end/handlers"
	services "pod-chef-back-end/internal/core/services"
	k8Repository "pod-chef-back-end/repositories/kubernetes"
	mongoRepository "pod-chef-back-end/repositories/mongodb"
)

func main() {
	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	kubernetesRepository := k8Repository.KubernetesRepository()
	nodeServices := services.NodeServices(kubernetesRepository.Nodes)
	podServices := services.PodServices(kubernetesRepository.Pods)
	namespaceServices := services.NamespaceServices(kubernetesRepository.Namespaces)
	deploymentServices := services.DeploymentServices(
		kubernetesRepository.Deployments,
		kubernetesRepository.Namespaces,
		kubernetesRepository.Services)
	serviceServices := services.ServiceServices(kubernetesRepository.Services)
	volumeServices := services.VolumeServices(kubernetesRepository.Volumes)

	mRepository := mongoRepository.MongoRepository()
	userServices := services.UserServices(mRepository.User)

	handlers.NodeHandler(e, nodeServices)
	handlers.PodHandler(e, podServices)
	handlers.DeploymentHandler(e, deploymentServices)
	handlers.NamespaceHandler(e, namespaceServices)
	handlers.ServiceHandler(e, serviceServices)
	handlers.VolumeHandler(e, volumeServices)
	handlers.AuthHandler(e, userServices)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
