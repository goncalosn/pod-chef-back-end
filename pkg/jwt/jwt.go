package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func CreateToken(userId string) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	// TODO FLAGS
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["expires"] = time.Now().Add(time.Minute * 15).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte("super_secret"))
	if err != nil {
		return "", err
	}
	return token, nil
}
