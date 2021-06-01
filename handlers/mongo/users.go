package handlers

import (
	"crypto/sha256"
	"net/http"

	pkg "pod-chef-back-end/pkg"

	"github.com/labstack/echo/v4"
)

//lkogin verifies given email and password and returns a token
func (h *HTTPHandler) login(c echo.Context) error {
	//geting form data
	email := c.FormValue("email")
	password := c.FormValue("password")

	//checking data for empty values
	if email == "" || password == "" {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	//call driver adapter responsible for getting the user from the database
	user, err := h.mongoServices.GetUserByEmail(email)

	if err != nil {
		//type assertion of custom Error to default error
		mongoError := err.(*pkg.Error)

		//return the error sent by the service
		return c.JSON(mongoError.Code, mongoError)

	} else if user == nil { //user doesn't exist

		return c.JSON(http.StatusNotFound, "Not found")

	} else { //user exists, verify user's password with a hash of the one sent

		//hash sent password
		hash := sha256.Sum256([]byte(password))

		//compare hashes
		if user.Hash != hash {
			//wrong password
			return c.JSON(http.StatusNotFound, "Not found")
		}
	}

	token, err := pkg.GenerateToken(h.viper, user.Name, email, user.Role)
	if err != nil {
		//type assertion of custom Error to default error
		tokenError := err.(*pkg.Error)

		//return the error sent by the service
		return c.JSON(tokenError.Code, tokenError)

	}

	return c.JSONPretty(http.StatusOK, token, " ")
}

//signup creates an user and returns a token
func (h *HTTPHandler) signup(c echo.Context) error {
	//geting form data
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	//checking data for empty values
	if email == "" || password == "" || name == "" {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	//call driver adapter responsible for getting the user from the database
	user, err := h.mongoServices.InsertUser(email, password, name, "user")

	if err != nil {
		//type assertion of custom Error to default error
		mongoError := err.(*pkg.Error)

		//return the error sent by the service
		return c.JSON(mongoError.Code, mongoError)

	}

	token, err := pkg.GenerateToken(h.viper, user.Name, email, user.Role)
	if err != nil {
		//type assertion of custom Error to default error
		tokenError := err.(*pkg.Error)

		//return the error sent by the service
		return c.JSON(tokenError.Code, tokenError)

	}

	return c.JSONPretty(http.StatusOK, token, " ")
}
