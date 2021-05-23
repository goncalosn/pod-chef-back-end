package user

import "go.mongodb.org/mongo-driver/mongo"

type MongoClient struct {
	Client *mongo.Client
}

func New(client *mongo.Client) *MongoClient {
	return &MongoClient{
		Client: client,
	}
}
