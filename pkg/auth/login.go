package auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func Login(email string, password string) {
	// procurar user na base de dados
	// auth user
	// create access token, write on header
}

func Signup(username string, email string, password string) {
	// generate password hash
	_, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Password := string(hashedPassword)
	// TokenHash := GenerateRandomString(16)
}
