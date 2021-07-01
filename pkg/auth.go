package pkg

import (
	"math/rand"
	"net/http"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

//GenerateJWT service responsible for generating a token
func GenerateJWT(viper *viper.Viper, name string, email string, role string, id string, date time.Time) (interface{}, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = name
	claims["email"] = email
	claims["role"] = role
	claims["id"] = id
	claims["date"] = date
	claims["exp"] = time.Now().UTC().Add(time.Hour * 72).Unix()

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

//EncryptPassword salts and hashes the password
func EncryptPassword(rawPassword string) []byte {
	hash, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil
	}
	return hash
}

//ComparePasswords compares the hashed password with a raw password
//
// Returns true if it matches
func ComparePasswords(hashedPassword string, rawPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(rawPassword))
	if err != nil {
		return false
	}
	return true
}

//GeneratePassword returns a string of 8 characters
func GeneratePassword() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	length := 8
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}
