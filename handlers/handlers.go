package http

import (
	deployments "pod-chef-back-end/handlers/deployments"
	nodes "pod-chef-back-end/handlers/nodes"
	pods "pod-chef-back-end/handlers/pods"
	ports "pod-chef-back-end/internal/core/ports"

	"github.com/labstack/echo/v4"
)

func PodHandler(e *echo.Echo, service ports.PodServices) {
	podsHandler := pods.NewHTTPHandler(service)

	e.GET("/pods", podsHandler.GetPodsByNodeAndNamespace)
}

func NodeHandler(e *echo.Echo, service ports.NodeServices) {
	nodesHandler := nodes.NewHTTPHandler(service)

	e.GET("/nodes", nodesHandler.GetNodes)
	e.GET("/node", nodesHandler.GetNode)
}

func DeploymentHandler(e *echo.Echo, service ports.DeploymentServices) {
	deploymentsHandler := deployments.NewHTTPHandler(service)

	e.GET("/deployments", deploymentsHandler.GetDeployments)
	e.DELETE("/deployment/:id", deploymentsHandler.DeleteDeployment)
	e.POST("/deployment/default-create", deploymentsHandler.CreateDefaultDeployment)
	e.POST("/deployment/advanced-create", deploymentsHandler.CreateFileDeployment)
}
