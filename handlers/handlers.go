package http

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	auths "pod-chef-back-end/handlers/auths"
	deployments "pod-chef-back-end/handlers/deployments"
	namespaces "pod-chef-back-end/handlers/namespaces"
	nodes "pod-chef-back-end/handlers/nodes"
	pods "pod-chef-back-end/handlers/pods"
	services "pod-chef-back-end/handlers/services"
	volumes "pod-chef-back-end/handlers/volumes"
	ports "pod-chef-back-end/internal/core/ports"

	"github.com/labstack/echo/v4"
)

var isLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("super_secret"),
})

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

func NamespaceHandler(e *echo.Echo, service ports.NamespaceServices) {
	namespacesHandler := namespaces.NewHTTPHandler(service)

	e.GET("/namespaces", namespacesHandler.GetNamespaces)
}

func ServiceHandler(e *echo.Echo, service ports.ServiceServices) {
	servicesHandler := services.NewHTTPHandler(service)

	e.GET("/services", servicesHandler.GetServicesByNamespace)
	e.GET("/service", servicesHandler.GetServiceByNameAndNamespace)
}

func VolumeHandler(e *echo.Echo, service ports.VolumeServices) {
	volumesHandler := volumes.NewHTTPHandler(service)

	e.GET("/volumes", volumesHandler.GetVolumes)
}

func AuthHandler(e *echo.Echo, service ports.UserServices) {
	authsHandler := auths.NewHTTPHandler(service)

	e.POST("/login", authsHandler.Login)
	e.POST("/signin", authsHandler.SignIn)
	e.GET("/private", private, isLoggedIn)
}

// TODO remove this, test purpose only
func private(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["user_id"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
