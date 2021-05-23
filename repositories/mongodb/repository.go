package mongodb

import (
	"pod-chef-back-end/repositories/mongodb/user"
)

type mongoRepository struct {
	User *user.MongoClient
}

func MongoRepository() *mongoRepository {
	client := Client()

	return &mongoRepository{
		User: user.New(client),
	}
}
