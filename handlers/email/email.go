package email

import (
	"net/http"
	"pod-chef-back-end/pkg"

	"github.com/labstack/echo/v4"
)

//newInvitationEmail create a new email with the
func (h *HTTPHandler) newInvitationEmail(c echo.Context) error {
	//getting form data
	receiver := c.FormValue("receiver")

	//checking data for empty values
	if receiver == "" {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	var subject = "You have been added to the Pod Chef whitelist!"

	//call driver adapter responsible for creating the deployment in the kubernetes cluster
	response, err := h.EmailServices.SendEmail(receiver, subject, "invitation.txt")
	if err != nil {
		//type assertion of custom Error to default error
		emailError := err.(*pkg.Error)

		//return the error sent by the service
		return c.JSON(emailError.Code, emailError)
	}

	return c.JSONPretty(http.StatusCreated, response, " ")
}

//newAnnulmentEmail create a new email with the
func (h *HTTPHandler) newAnnulmentEmail(c echo.Context) error {
	//getting form data
	receiver := c.FormValue("receiver")

	//checking data for empty values
	if receiver == "" {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	var subject = "You have been removed from the Pod Chef whitelist!"

	//call driver adapter responsible for creating the deployment in the kubernetes cluster
	response, err := h.EmailServices.SendEmail(receiver, subject, "annulment.txt")
	if err != nil {
		//type assertion of custom Error to default error
		emailError := err.(*pkg.Error)

		//return the error sent by the service
		return c.JSON(emailError.Code, emailError)
	}

	return c.JSONPretty(http.StatusCreated, response, " ")
}
