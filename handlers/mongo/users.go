package mongo

import (
	"net/http"
	"regexp"

	"pod-chef-back-end/pkg"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

//lkogin verifies given email and password and returns a token
func (h *HTTPHandler) login(c echo.Context) error {
	//body structure
	type body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	data := new(body)

	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	//checking data for empty values
	if data.Email == "" || data.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	//call driver adapter responsible for getting the user from the database
	user, err := h.mongoServices.GetUserByEmail(data.Email)

	if err != nil {
		//type assertion of custom Error to default error
		mongoError := err.(*pkg.Error)

		//return the error sent by the service
		return c.JSON(mongoError.Code, mongoError)

	} else if user == nil { //user doesn't exist
		return c.JSON(http.StatusNotFound, map[string]string{"message": "Not found"})
	} else { //user exists, verify user's password with a hash of the one sent
		//compare hashes
		if !pkg.ComparePasswords(user.Hash, data.Password) {
			//wrong password
			return c.JSON(http.StatusNotFound, map[string]string{"message": "Not found"})
		}
	}

	token, err := pkg.GenerateJWT(h.viper, user.Name, user.Email, user.Role)
	if err != nil {
		//type assertion of custom Error to default error
		tokenError := err.(*pkg.Error)

		//return the error sent by the service
		return c.JSON(tokenError.Code, tokenError)
	}

	return c.JSON(http.StatusOK, token)
}

//signup creates an user and returns a token
func (h *HTTPHandler) signup(c echo.Context) error {
	//body structure
	type body struct {
		Email    string `json:"email"`
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	data := new(body)

	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	//checking data for empty values
	if data.Email == "" || data.Password == "" || data.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	//verify email lenght
	if len(data.Email) < 3 && len(data.Email) > 254 {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Email not valid"})
	}

	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	//validate email
	if !emailRegex.MatchString(data.Email) {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Email not valid"})
	}

	//verify password length
	if len(data.Password) < 7 {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Password requires a minimum of 7 characters"})
	}

	crypt := pkg.EncryptPassword(data.Password)
	//call driver adapter responsible for getting the user from the database
	user, err := h.mongoServices.InsertUser(data.Email, string(crypt), data.Name)

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
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	//call driver adapter responsible for deleting a user from the database
	response, err := h.mongoServices.DeleteUser(data.Email)

	if err != nil {
		//type assertion of custom Error to default error
		mongoError := err.(*pkg.Error)

		//return the error sent by the service
		return c.JSON(mongoError.Code, mongoError)
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}

//updateSelfPassword update password
func (h *HTTPHandler) updateSelfPassword(c echo.Context) error {
	//body structure
	type body struct {
		Password string `json:"password"`
	}

	data := new(body)

	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	//checking data for empty values
	if len(data.Password) < 7 {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Password requires a minimum of 7 characters"})
	}

	//get the token's claims
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	crypt := pkg.EncryptPassword(data.Password)
	//call driver adapter responsible for updating a user's password from the database
	response, err := h.mongoServices.UpdateSelfPassword(email, string(crypt))

	if err != nil {
		//type assertion of custom Error to default error
		mongoError := err.(*pkg.Error)

		//return the error sent by the service
		return c.JSON(mongoError.Code, mongoError)
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}

//resetPassword update password
func (h *HTTPHandler) resetPassword(c echo.Context) error {
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

	//generate not so random password
	generated := pkg.GeneratePassword()

	//call driver adapter responsible for reseting a user's password
	response, err := h.mongoServices.ResetUserPassword(data.Email, generated)

	if err != nil {
		//type assertion of custom Error to default error
		mongoError := err.(*pkg.Error)

		//return the error sent by the service
		return c.JSON(mongoError.Code, mongoError)
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}

//updateSelfName update name
func (h *HTTPHandler) updateSelfName(c echo.Context) error {
	//body structure
	type body struct {
		Name string `json:"name"`
	}

	data := new(body)

	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	//checking data for empty values
	if data.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	//get the token's claims
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	email := claims["email"].(string)

	//call driver adapter responsible for updating a user's name from the database
	response, err := h.mongoServices.UpdateUserName(email, data.Name)

	if err != nil {
		//type assertion of custom Error to default error
		mongoError := err.(*pkg.Error)

		//return the error sent by the service
		return c.JSON(mongoError.Code, mongoError)
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}

//updateUserRole update user role
func (h *HTTPHandler) updateUserRole(c echo.Context) error {
	//body structure
	type body struct {
		Email string `json:"email"`
		Role  string `json:"role"`
	}

	data := new(body)

	if err := c.Bind(data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}
	//checking data for empty values
	if data.Email == "" || data.Role == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	//checking for role validity
	if data.Role != "member" && data.Role != "admin" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	//call driver adapter responsible updating a user's role
	response, err := h.mongoServices.UpdateUserRole(data.Email, data.Role)

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
	response, err := h.mongoServices.InviteUserToWhitelist(data.Email)
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
	response, err := h.mongoServices.RemoveUserFromWhitelist(data.Email)
	if err != nil {
		//type assertion of custom Error to default error
		emailError := err.(*pkg.Error)

		//return the error sent by the service
		return c.JSON(emailError.Code, emailError)
	}

	return c.JSONPretty(http.StatusOK, response, " ")
}
