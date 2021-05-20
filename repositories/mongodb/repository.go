package mongodb

import (
	users "pod-chef-back-end/repositories/mongodb/users"
)

type mongoRepository struct {
	User *users.MongoClient
}

func MongoRepository() *mongoRepository {
	client := Client()

	return &mongoRepository{
		User: users.New(client),
	}
}
