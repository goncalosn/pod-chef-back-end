package handlers

import (
	ports "pod-chef-back-end/internal/core/ports"

	echo "github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

//HTTPHandler mongo services
type HTTPHandler struct {
	mongoServices ports.MongoServices
	viper         *viper.Viper
}

//NewHTTPHandler where services are injected
func NewHTTPHandler(mongoServices ports.MongoServices, viper *viper.Viper) *HTTPHandler {
	return &HTTPHandler{
		mongoServices: mongoServices,
		viper:         viper,
	}
}

//Handlers contains containers every handler associated with kubernetes
func Handlers(e *echo.Echo, service *HTTPHandler, isLoggedIn echo.MiddlewareFunc) {
	e.POST("/login", service.login)
	e.POST("/signup", service.signup)
}
