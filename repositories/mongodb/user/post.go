package user

import (
	"context"
	"errors"
	"fmt"
	"pod-chef-back-end/pkg/auth"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (serviceHandler *MongoClient) Register(username string, email string, rawPassword string) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	token := auth.GenerateTokenHash()
	fmt.Println([]byte(token))
	// TODO this throws errors???
	password := auth.EncryptPassword(rawPassword, token)
	fmt.Println([]byte(password))

	// TODO dont save empty users
	res, err := serviceHandler.Client.Database("main").Collection("users").InsertOne(ctx, bson.M{
		"username": username,
		"email":    email,
		"password": password,
		"token":    token,
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
	if err := serviceHandler.Client.Database("main").Collection("users").FindOne(ctx, bson.M{"email": email}).Decode(&user); err != nil {
		return nil, err
	}
	if auth.DecryptPassword(user.Password, user.Token) != rawPassword {
		return nil, errors.New("wrong password")
	}

	return user, nil
}
