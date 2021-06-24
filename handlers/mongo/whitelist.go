package mongo

import (
	"net/http"
	"pod-chef-back-end/pkg"

	echo "github.com/labstack/echo/v4"
)

//getAllUsersFromWhitelist get all the users from the whitelist
func (h *HTTPHandler) getAllUsersFromWhitelist(c echo.Context) error {
	//call driver adapter responsible for getting all the users from the database
	response, err := h.mongoServices.GetAllUsersFromWhitelist()

	if err != nil {
		//type assertion of custom Error to default error
		mongoError := err.(*pkg.Error)

		//return the error sent by the service
		return c.JSON(mongoError.Code, mongoError)
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}

//inviteUserToWhitelist create a new email with the
func (h *HTTPHandler) inviteUserToWhitelist(c echo.Context) error {
	//body structure
	type body struct {
		Email string `json:"email"`
	}

	data := new(body)

	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}
	//checking data for empty values
	if data.Email == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	//call driver adapter responsible for creating the deployment in the kubernetes cluster
	response, err := h.mongoServices.InsertUserIntoWhitelist(data.Email)
	if err != nil {
		//type assertion of custom Error to default error
		emailError := err.(*pkg.Error)

		//return the error sent by the service
		return c.JSON(emailError.Code, emailError)
	}

	return c.JSONPretty(http.StatusCreated, response, " ")
}

//removeUserFromWhitelist create a new email with the
func (h *HTTPHandler) removeUserFromWhitelist(c echo.Context) error {
	//body structure
	type body struct {
		ID string `json:"id"`
	}

	data := new(body)

	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	//checking data for empty values
	if data.ID == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	//call driver adapter responsible for creating the deployment in the kubernetes cluster
	response, err := h.mongoServices.RemoveUserFromWhitelist(data.ID)
	if err != nil {
		//type assertion of custom Error to default error
		emailError := err.(*pkg.Error)

		//return the error sent by the service
		return c.JSON(emailError.Code, emailError)
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}
