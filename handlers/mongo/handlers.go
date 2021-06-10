package mongo

import (
	"pod-chef-back-end/internal/core/ports"
	"pod-chef-back-end/pkg"

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
func Handlers(e *echo.Echo, handler *HTTPHandler, isLoggedIn echo.MiddlewareFunc) {
	e.POST("/login", handler.login)
	e.POST("/signup", handler.signup)
	e.GET("/users", handler.getAllUsers)
	e.DELETE("/user", handler.deleteUser)
	e.GET("/whitelist", handler.getAllUsersFromWhitelist)
	e.POST("/whitelist", handler.inviteUserToWhitelist, isLoggedIn, pkg.IsAdmin)
	e.DELETE("/whitelist", handler.removeUserFromWhitelist, isLoggedIn, pkg.IsAdmin)
}
