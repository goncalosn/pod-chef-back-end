package auth

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
)

type User struct {
	ID       string `json:"_id" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Token    string `json:"token" bson:"token"`
}

var (
	key = []byte("passphrasewhichneedstobe32bytes!")
)

func GenerateTokenHash() []byte {
	token := make([]byte, aes.BlockSize)
	if _, err := rand.Read(token); err != nil {
		return []byte("1234567890123456")
	}
	return token
}

func EncryptPassword(rawPassword string, iv []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key error1", err)
	}

	ecb := cipher.NewCBCEncrypter(block, iv)
	content := []byte(rawPassword)
	content = PKCS5Padding(content)
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)

	return crypted
}

func DecryptPassword(password []byte, iv []byte) string {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key error1", err)
	}

	ecb := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(password))
	ecb.CryptBlocks(decrypted, password)

	return string(PKCS5Trimming(decrypted))
}

func PKCS5Padding(ciphertext []byte) []byte {
	padding := aes.BlockSize - len(ciphertext)%aes.BlockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
