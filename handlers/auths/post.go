package auths

import (
	"encoding/json"
	"net/http"
	"pod-chef-back-end/pkg/auth"

	"github.com/labstack/echo/v4"
)

func (h *HTTPHandler) Login(c echo.Context) error {
	email := c.FormValue("email")
	//TODO this shit, make login properly
	res, err := h.UserServices.Login(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Login Failed")
	}

	return c.JSONPretty(http.StatusCreated, res, " ")
}

func (h *HTTPHandler) SignIn(c echo.Context) error {
	var user auth.User
	if err := json.NewDecoder(c.Request().Body).Decode(&user); err != nil {
		c.JSON(http.StatusInternalServerError, "Error parsing json")
	}
	res, err := h.UserServices.SignIn(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Login Failed")
	}

	return c.JSONPretty(http.StatusCreated, res, " ")
}
