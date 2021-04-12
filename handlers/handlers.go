package http

import (
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
