package pkg

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

//GenerateToken service responsible for generating a token
func GenerateToken(viper *viper.Viper, name string, email string, role string) (interface{}, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = name
	claims["email"] = name
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	secret := viper.Get("TOKEN_SECRET").(string)

	// Generate encoded token and send it as response.
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		//print the error stack
		log.Error(err)

		//return a custom error
		return nil, &Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}
	return map[string]string{"token": signedToken}, nil
}

//IsAdmin verifies the token's claims and checks the user role to verify if it is admin
func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		role := claims["role"].(string)
		if role != "admin" {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}
