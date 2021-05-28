package http

import (
	ports "pod-chef-back-end/internal/core/ports"

	echo "github.com/labstack/echo/v4"
)

//HTTPHandler kubernetes services
type HTTPHandler struct {
	kubernetesServices ports.KubernetesServices
}

//NewHTTPHandler where the kubernetes services are injected
func NewHTTPHandler(kubernetesServices ports.KubernetesServices) *HTTPHandler {
	return &HTTPHandler{
		kubernetesServices: kubernetesServices,
	}
}

//NodesHandler containers every handler associated with nodes
func NodesHandler(e *echo.Echo, service *HTTPHandler) {
	e.GET("/nodes", service.GetNodes)
	e.GET("/node", service.GetNodeByName)
}

//DeploymentsHandler containers every handler associated with deployments
func DeploymentsHandler(e *echo.Echo, service *HTTPHandler) {
	e.GET("/deployments", service.GetDeploymentsByUserAndName)
	e.GET("/deployment", service.GetDeploymentsByUser)
	e.POST("/deployment", service.CreateDeployment)
}

//NamespacesHandler containers every handler associated with namespaces
func NamespacesHandler(e *echo.Echo, service *HTTPHandler) {
	e.GET("/namespaces", service.GetNamespaces)
}
