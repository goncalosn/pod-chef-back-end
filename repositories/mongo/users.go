package services

import (
	"context"
	"net/http"

	models "pod-chef-back-end/internal/core/domain/mongo"
	"pod-chef-back-end/pkg"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//GetUserByEmail method responsible for getting a user from the database
func (repo *MongoRepository) GetUserByEmail(email string) (*models.User, error) {
	//data structure to where the data will be written to
	var user *models.User

	//choose the database and collection
	collection := repo.Client.Database("podchef").Collection("users")

	//data to filter the search with
	filter := bson.M{"email": email}

	//call driven adapter responsible for getting a user data from the database
	err := collection.FindOne(context.Background(), filter).Decode(&user)

	if err != nil && err != mongo.ErrNoDocuments {
		//print the error stack
		log.Error(err)

		//return a custom error
		return nil, &pkg.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	return user, nil
}

//InsertUser method responsible for inserting a user in the database
func (repo *MongoRepository) InsertUser(email string, hash string, name string, role string) (*models.User, error) {
	//data structure containing the data to be inserted
	user := &models.User{
		Email:    email,
		Password: hash,
		Name:     name,
		Role:     role,
	}

	//choose the database and collection
	collection := repo.Client.Database("podchef").Collection("users")

	//call driven adapter responsible for inserting a user into the database
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		//print the error stack
		log.Error(err)

		//return a custom error
		return nil, &pkg.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	return user, nil
}

//DeleteUserByEmail method responsible for deleting a user
func (repo *MongoRepository) DeleteUserByEmail(email string) (interface{}, error) {
	//choose the database and collection
	collection := repo.Client.Database("podchef").Collection("users")

	//data to filter with
	filter := bson.D{{"email", email}}

	//call driven adapter responsible for deleting a user from the database
	response, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		//print the error stack
		log.Error(err)

		//return a custom error
		return nil, &pkg.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	return response, nil
}
