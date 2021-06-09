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
func NewHTTPHandler(kubernetesServices ports.KubernetesServices) *HTTPHandler {
	return &HTTPHandler{
		kubernetesServices: kubernetesServices,
	}
}

//Handlers contains containers every handler associated with kubernetes
func Handlers(e *echo.Echo, handler *HTTPHandler, isLoggedIn echo.MiddlewareFunc) {
	e.GET("/nodes", handler.getNodes, isLoggedIn, pkg.IsAdmin)
	e.GET("/node", handler.getNodeByName, isLoggedIn, pkg.IsAdmin)

	e.POST("/deployment", handler.createDeployment, isLoggedIn)
	e.GET("/deployments", handler.getDeploymentsByUser, isLoggedIn)
	e.GET("/deployment", handler.getDeploymentByUserAndName, isLoggedIn)
	e.DELETE("/deployment", handler.deleteDeploymentByUserAndName, isLoggedIn)
}
