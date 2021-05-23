package auth

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

type User struct {
	ID       string `json:"id" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Token    string `json:"token" bson:"token"`
}

func GenerateTokenHash() string {
	token := make([]byte, aes.BlockSize)
	if _, err := rand.Read(token); err != nil {
		return "1234567890123456"
	}
	return string(token)
}

func EncryptPassword(rawPassword string, iv string) string {
	bRawPassword := []byte(rawPassword)
	passLen := len(rawPassword)
	bIv := []byte(iv)
	key := []byte("passphrasewhichneedstobe32bytes!")

	// generate a new aes cipher using our 32 byte long key
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	if passLen%aes.BlockSize != 0 {
		bRawPassword = PKCS5Padding(bRawPassword, aes.BlockSize)
	}
	ciphertext := make([]byte, aes.BlockSize+len(bRawPassword))
	mode := cipher.NewCBCEncrypter(block, bIv)
	mode.CryptBlocks(ciphertext, bRawPassword)

	return string(ciphertext)
}

func DecryptPassword(password string, iv string) string {
	cyphertext, _ := hex.DecodeString(password)
	bIv := []byte(iv)
	key := []byte("passphrasewhichneedstobe32bytes!")

	// generate a new aes cipher using our 32 byte long key
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	mode := cipher.NewCBCDecrypter(block, bIv)
	mode.CryptBlocks(cyphertext, cyphertext)

	return string(PKCS5Trimming(cyphertext))
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
