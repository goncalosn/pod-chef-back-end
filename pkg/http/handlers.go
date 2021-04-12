package http

import (
	"net/http"

	builders "pod-chef-back-end/pkg/builders"
	nodes "pod-chef-back-end/pkg/http/nodes"
	pods "pod-chef-back-end/pkg/http/pods"

	"github.com/labstack/echo/v4"
)

func BuildHandlers(e *echo.Echo, interactors *builders.InteractorsContainer) {

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	podHandler := pods.PodHandler{PodInteractor: *interactors.PodInteractor}
	nodeHandler := nodes.NodeHandler{NodeInteractor: *interactors.NodeInteractor}

	e.GET("/pods", podHandler.GetPodsByNodeAndNamespace)
	e.GET("/nodes", nodeHandler.GetNodes)
	e.GET("/node", nodeHandler.GetNode)
}
