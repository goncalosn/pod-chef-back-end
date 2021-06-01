package kubernetes

import (
	ports "pod-chef-back-end/internal/core/ports"
	"pod-chef-back-end/pkg"

	echo "github.com/labstack/echo/v4"
)

//HTTPHandler kubernetes services
type HTTPHandler struct {
	kubernetesServices ports.KubernetesServices
	mongoServices      ports.MongoServices
}

//NewHTTPHandler where services are injected
func NewHTTPHandler(kubernetesServices ports.KubernetesServices, mongoServices ports.MongoServices) *HTTPHandler {
	return &HTTPHandler{
		kubernetesServices: kubernetesServices,
		mongoServices:      mongoServices,
	}
}

//Handlers contains containers every handler associated with kubernetes
func Handlers(e *echo.Echo, service *HTTPHandler, isLoggedIn echo.MiddlewareFunc) {
	e.GET("/nodes", service.getNodes, isLoggedIn, pkg.IsAdmin)
	e.GET("/node", service.getNodeByName, isLoggedIn, pkg.IsAdmin)

	e.GET("/deployments", service.getDeploymentsByUserAndName, isLoggedIn)
	e.GET("/deployment", service.getDeploymentsByUser, isLoggedIn)
	e.POST("/deployment", service.createDeployment, isLoggedIn)
}
