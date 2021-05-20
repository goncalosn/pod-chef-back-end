package users

import (
	"context"
	"log"
	"pod-chef-back-end/pkg/auth"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (serviceHandler *MongoClient) SignIn(user auth.User) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := serviceHandler.Client.Database("Default").Collection("users").InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return res, nil
}

func (serviceHandler *MongoClient) Login(email string) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var user auth.User
	if err := serviceHandler.Client.Database("Default").Collection("users").FindOne(ctx, bson.M{"email": email}).Decode(&user); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return user, nil
}
