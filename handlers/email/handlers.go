package email

import (
	"pod-chef-back-end/internal/core/ports"
	"pod-chef-back-end/pkg"

	echo "github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

//HTTPHandler email services
type HTTPHandler struct {
	EmailServices ports.EmailServices
	viper         *viper.Viper
}

//NewHTTPHandler where services are injected
func NewHTTPHandler(emailServices ports.EmailServices, viper *viper.Viper) *HTTPHandler {
	return &HTTPHandler{
		EmailServices: emailServices,
		viper:         viper,
	}
}

//Handlers contains containers every handler associated with kubernetes
func Handlers(e *echo.Echo, service *HTTPHandler, isLoggedIn echo.MiddlewareFunc) {
	e.POST("/email/invitation", service.newInvitationEmail, isLoggedIn, pkg.IsAdmin)
	e.POST("/email/annulment", service.newAnnulmentEmail, isLoggedIn, pkg.IsAdmin)
}
