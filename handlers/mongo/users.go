package handlers

import (
	"encoding/base64"
	"encoding/hex"
	"net/http"
	"pod-chef-back-end/internal/core/domain/mongo"

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
	if reqUser.Email == "" || reqUser.Password == "" {
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
		// generate token IV for aes-cbc encryption
		tokenIv, err := base64.StdEncoding.DecodeString(user.TokenIv)
		if err != nil {
			//type assertion of custom Error to default error
			hb64DecodeError := err.(*pkg.Error)
			//return the error sent by the service
			return c.JSON(hb64DecodeError.Code, hb64DecodeError)
		}
		// get key in a hex string
		hexKey := h.viper.Get("CRYPT_KEY").(string)
		// convert hex string to a byte array
		key, err := hex.DecodeString(hexKey)
		if err != nil {
			//type assertion of custom Error to default error
			hexDecodeError := err.(*pkg.Error)
			//return the error sent by the service
			return c.JSON(hexDecodeError.Code, hexDecodeError)
		}
		crypt := pkg.EncryptPassword(reqUser.Password, tokenIv, key)
		//compare hashes
		if user.Password != base64.StdEncoding.EncodeToString(crypt) {
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
	if reqUser.Email == "" || reqUser.Password == "" || reqUser.Name == "" {
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}
	// generate token IV for aes-cbc encryption
	tokenIv := pkg.GenerateTokenIV()
	// get key in a hex string
	hexKey := h.viper.Get("CRYPT_KEY").(string)
	// convert hex string to a byte array
	key, err := hex.DecodeString(hexKey)
	if err != nil {
		//type assertion of custom Error to default error
		hexDecodeError := err.(*pkg.Error)
		//return the error sent by the service
		return c.JSON(hexDecodeError.Code, hexDecodeError)
	}
	crypt := pkg.EncryptPassword(reqUser.Password, tokenIv, key)
	//call driver adapter responsible for getting the user from the database
	user, err := h.mongoServices.InsertUser(reqUser.Email, base64.StdEncoding.EncodeToString(crypt), base64.StdEncoding.EncodeToString(tokenIv), reqUser.Name, "user")

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
