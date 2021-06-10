package mongo

import (
	"net/http"
	"pod-chef-back-end/internal/core/domain/mongo"
	"regexp"

	"pod-chef-back-end/pkg"

	"github.com/labstack/echo/v4"
)

//lkogin verifies given email and password and returns a token
func (h *HTTPHandler) login(c echo.Context) error {
	// request body in json, bind request json to struct
	reqUser := &mongo.User{}
	err := c.Bind(reqUser)
	if err != nil {
		//type assertion of custom Error to default error
		JSONBindError := err.(*pkg.Error)
		//return the error sent by the service
		return c.JSON(JSONBindError.Code, JSONBindError)
	}

	//checking data for empty values
	if reqUser.Email == "" || reqUser.Hash == "" {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	//call driver adapter responsible for getting the user from the database
	user, err := h.mongoServices.GetUserByEmail(reqUser.Email)

	if err != nil {
		//type assertion of custom Error to default error
		mongoError := err.(*pkg.Error)

		//return the error sent by the service
		return c.JSON(mongoError.Code, mongoError)

	} else if user == nil { //user doesn't exist
		return c.JSON(http.StatusNotFound, "Not found")
	} else { //user exists, verify user's password with a hash of the one sent
		//compare hashes
		if !pkg.ComparePasswords(user.Hash, reqUser.Hash) {
			//wrong password
			return c.JSON(http.StatusNotFound, "Not found")
		}
	}

	token, err := pkg.GenerateJWT(h.viper, user.Name, user.Email, user.Role)
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
	// request body in json, bind request json to struct
	reqUser := &mongo.User{}
	err := c.Bind(reqUser)
	if err != nil {
		//type assertion of custom Error to default error
		JSONBindError := err.(*pkg.Error)
		//return the error sent by the service
		return c.JSON(JSONBindError.Code, JSONBindError)
	}

	//checking data for empty values
	if reqUser.Email == "" || reqUser.Hash == "" || reqUser.Name == "" {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	//verify email lenght
	if len(reqUser.Email) < 3 && len(reqUser.Email) > 254 {
		return c.JSON(http.StatusBadRequest, "Email not valid")
	}

	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	//validate email
	if !emailRegex.MatchString(reqUser.Email) {
		return c.JSON(http.StatusBadRequest, "Email not valid")
	}

	//verify password length
	if len(reqUser.Hash) < 7 {
		return c.JSON(http.StatusBadRequest, "Password requires a minimum of 7 characters")
	}

	crypt := pkg.EncryptPassword(reqUser.Hash)
	//call driver adapter responsible for getting the user from the database
	user, err := h.mongoServices.InsertUser(reqUser.Email, string(crypt), reqUser.Name, "user")

	if err != nil {
		//type assertion of custom Error to default error
		mongoError := err.(*pkg.Error)

		//return the error sent by the service
		return c.JSON(mongoError.Code, mongoError)
	}
	// generate json web token
	token, err := pkg.GenerateJWT(h.viper, user.Name, user.Email, user.Role)
	if err != nil {
		//type assertion of custom Error to default error
		tokenError := err.(*pkg.Error)

		//return the error sent by the service
		return c.JSON(tokenError.Code, tokenError)
	}

	return c.JSONPretty(http.StatusOK, token, " ")
}

//getAllUsers get all the users from the database
func (h *HTTPHandler) getAllUsers(c echo.Context) error {
	//call driver adapter responsible for getting all the users from the database
	response, err := h.mongoServices.GetAllUsers()

	if err != nil {
		//type assertion of custom Error to default error
		mongoError := err.(*pkg.Error)

		//return the error sent by the service
		return c.JSON(mongoError.Code, mongoError)
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}

//deleteUser delete user ffrom the system
func (h *HTTPHandler) deleteUser(c echo.Context) error {
	//geting query data
	email := c.QueryParam("email")

	//checking data for empty values
	if email == "" {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	//call driver adapter responsible for deleting a user from the database
	response, err := h.mongoServices.DeleteUser(email)

	if err != nil {
		//type assertion of custom Error to default error
		mongoError := err.(*pkg.Error)

		//return the error sent by the service
		return c.JSON(mongoError.Code, mongoError)
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}

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
	//getting form data
	email := c.FormValue("email")

	//checking data for empty values
	if email == "" {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	//call driver adapter responsible for creating the deployment in the kubernetes cluster
	response, err := h.mongoServices.InviteUserToWhitelist(email)
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
	//getting form data
	email := c.FormValue("email")

	//checking data for empty values
	if email == "" {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	//call driver adapter responsible for creating the deployment in the kubernetes cluster
	response, err := h.mongoServices.RemoveUserFromWhitelist(email)
	if err != nil {
		//type assertion of custom Error to default error
		emailError := err.(*pkg.Error)

		//return the error sent by the service
		return c.JSON(emailError.Code, emailError)
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}
