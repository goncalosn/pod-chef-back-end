package auths

import (
	"encoding/json"
	"net/http"
	"pod-chef-back-end/pkg/auth"

	"github.com/labstack/echo/v4"
)

func (h *HTTPHandler) Login(c echo.Context) error {
	var user auth.User
	if err := json.NewDecoder(c.Request().Body).Decode(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, "Error parsing json")
	}
	res, err := h.UserServices.Authenticate(user.Email, user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Authenticate Failed")
	}

	return c.JSONPretty(http.StatusCreated, res, " ")
}

func (h *HTTPHandler) SignIn(c echo.Context) error {
	var user auth.User
	if err := json.NewDecoder(c.Request().Body).Decode(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, "Error parsing json")
	}
	res, err := h.UserServices.Register(user.Username, user.Email, user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Authentication Error")
	}

	return c.JSONPretty(http.StatusCreated, res, " ")
}
