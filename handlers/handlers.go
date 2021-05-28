package http

import (
	ports "pod-chef-back-end/internal/core/ports"

	echo "github.com/labstack/echo/v4"
)

type HTTPHandler struct {
	KubernetesServices ports.KubernetesServices
}

func NewHTTPHandler(kubernetesServices ports.KubernetesServices) *HTTPHandler {
	return &HTTPHandler{
		KubernetesServices: kubernetesServices,
	}
}

func NodesHandler(e *echo.Echo, service *HTTPHandler) {
	e.GET("/nodes", service.GetNodes)
	e.GET("/node", service.GetNode)
}

func DeploymentsHandler(e *echo.Echo, service *HTTPHandler) {
	e.GET("/deployments", service.GetDeploymentsByNamespace)
	e.DELETE("/deployment/:id", service.DeleteDeployment)
	e.POST("/deployment", service.CreateDeployment)
}

func NamespacesHandler(e *echo.Echo, service *HTTPHandler) {
	e.GET("/namespaces", service.GetNamespaces)
}
