package user

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"pod-chef-back-end/pkg/auth"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (serviceHandler *MongoClient) Register(username string, email string, rawPassword string) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	token := auth.GenerateTokenHash()
	password := auth.EncryptPassword(rawPassword, token)

	// TODO dont save empty users
	res, err := serviceHandler.Client.Database("main").Collection("users").InsertOne(ctx, bson.M{
		"username": username,
		"email":    email,
		"password": base64.StdEncoding.EncodeToString(password),
		"token":    base64.StdEncoding.EncodeToString(token),
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (serviceHandler *MongoClient) Authenticate(email string, rawPassword string) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user auth.User
	err := serviceHandler.Client.Database("main").Collection("users").FindOne(ctx, bson.D{{"email", email}}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("no email in db")
		}
		return nil, err
	}
	token, _ := base64.StdEncoding.DecodeString(user.Token)
	crypt, _ := base64.StdEncoding.DecodeString(user.Password)
	decrypt := auth.DecryptPassword(crypt, token)
	if decrypt != rawPassword {
		return nil, errors.New("wrong password")
	}

	return user, nil
}
