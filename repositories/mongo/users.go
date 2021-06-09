package services

import (
	"context"
	"net/http"
	"time"

	models "pod-chef-back-end/internal/core/domain/mongo"
	"pod-chef-back-end/pkg"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, &pkg.Error{Err: err, Code: http.StatusNotFound, Message: "User not found"}
		}
		//print the error stack
		log.Error(err)

		//return a custom error
		return nil, &pkg.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	return user, nil
}

//GetAllUsers method responsible for all users
func (repo *MongoRepository) GetAllUsers() (interface{}, error) {
	//data structure to where the data will be written to
	var users []bson.M

	//choose the database and collection
	collection := repo.Client.Database("podchef").Collection("users")

	//data to filter the search with
	filter := bson.D{}

	//call driven adapter responsible for getting a users's data from the database to a cursor
	cur, err := collection.Find(context.Background(), filter, &options.FindOptions{Projection: bson.M{"_id": 0, "hash": 0}})

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, &pkg.Error{Err: err, Code: http.StatusNotFound, Message: "No user were found"}
		}
		//print the error stack
		log.Error(err)

		//return a custom error
		return nil, &pkg.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	//decode all the data from the cursor into the users's structure
	err = cur.All(context.Background(), &users)
	if err != nil {
		//print the error stack
		log.Error(err)

		//return a custom error
		return nil, &pkg.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	return users, nil
}

//InsertUser method responsible for inserting a user in the database
func (repo *MongoRepository) InsertUser(email string, hash string, name string, role string) (*models.User, error) {
	//data structure containing the data to be inserted
	user := &models.User{
		Email: email,
		Hash:  hash,
		Name:  name,
		Role:  role,
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

//GetUserFromWhitelistByEmail method responsible for getting a user from the database
func (repo *MongoRepository) GetUserFromWhitelistByEmail(email string) (interface{}, error) {
	//data structure to where the data will be written to
	var user interface{}

	//choose the database and collection
	collection := repo.Client.Database("podchef").Collection("whitelist")

	//data to filter the search with
	filter := bson.M{"email": email}

	//call driven adapter responsible for getting a user data from the database
	err := collection.FindOne(context.Background(), filter).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, &pkg.Error{Err: err, Code: http.StatusNotFound, Message: "User not found"}
		}
		//print the error stack
		log.Error(err)

		//return a custom error
		return nil, &pkg.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	return user, nil
}

//GetAllUsersFromWhitelist method responsible for all users fro mthe whitelist
func (repo *MongoRepository) GetAllUsersFromWhitelist() (interface{}, error) {
	//data structure to where the data will be written to
	var users []bson.M

	//choose the database and collection
	collection := repo.Client.Database("podchef").Collection("whitelist")

	//data to filter the search with
	filter := bson.D{}

	//call driven adapter responsible for getting a users's data from the database to a cursor
	cur, err := collection.Find(context.Background(), filter, &options.FindOptions{Projection: bson.M{"_id": 0}})

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, &pkg.Error{Err: err, Code: http.StatusNotFound, Message: "No users were found"}
		}
		//print the error stack
		log.Error(err)

		//return a custom error
		return nil, &pkg.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	//decode all the data from the cursor into the users's structure
	err = cur.All(context.Background(), &users)

	if err != nil {
		//print the error stack
		log.Error(err)

		//return a custom error
		return nil, &pkg.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	return users, nil
}

//InsertUserIntoWhitelist method responsible for inserting a user in the database
func (repo *MongoRepository) InsertUserIntoWhitelist(email string) (bool, error) {
	//data structure containing the data to be inserted
	user := struct {
		Email string    `bson:"email"`
		Data  time.Time `bson:"data"`
	}{
		Email: email,
		Data:  time.Now().UTC(),
	}

	//choose the database and collection
	collection := repo.Client.Database("podchef").Collection("whitelist")

	//call driven adapter responsible for inserting a user into the database
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		//print the error stack
		log.Error(err)

		//return a custom error
		return false, &pkg.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	return true, nil
}

//DeleteUserFromWhitelistByEmail method responsible for deleting a user
func (repo *MongoRepository) DeleteUserFromWhitelistByEmail(email string) (bool, error) {
	//choose the database and collection
	collection := repo.Client.Database("podchef").Collection("whitelist")

	//data to filter with
	filter := bson.D{{"email", email}}

	//call driven adapter responsible for deleting a user from the database
	_, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, &pkg.Error{Err: err, Code: http.StatusNotFound, Message: "User not found"}
		}
		//print the error stack
		log.Error(err)

		//return a custom error
		return false, &pkg.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	return true, nil
}
