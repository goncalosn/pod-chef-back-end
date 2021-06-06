package pkg

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

//GenerateJWT service responsible for generating a token
func GenerateJWT(viper *viper.Viper, name string, email string, role string) (interface{}, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = name
	claims["email"] = email
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

//GenerateTokenIV generates random bytes into a buffer, this creates the IV for the user
func GenerateTokenIV() []byte {
	token := make([]byte, aes.BlockSize)
	if _, err := rand.Read(token); err != nil {
		log.Fatal(err)
	}
	return token
}

//EncryptPassword encrypts plaintext password into AES-CBC, return a byte array
func EncryptPassword(rawPassword string, iv []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key error1", err)
	}

	// cbc block
	ecb := cipher.NewCBCEncrypter(block, iv)
	content := []byte(rawPassword)
	content = PKCS5Padding(content)
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)

	return crypted
}

//PKCS5Padding pads a byte array to 16 multiples (%16 == 0)
func PKCS5Padding(ciphertext []byte) []byte {
	padding := aes.BlockSize - len(ciphertext)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
