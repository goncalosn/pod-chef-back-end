package mongo

import (
	"context"
	"net/http"
	models "pod-chef-back-end/internal/core/domain/mongo"
	pkg "pod-chef-back-end/pkg"
	"time"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//GetUserFromWhitelistByEmail method responsible for getting a user from the database
func (repo *MongoRepository) GetUserFromWhitelistByEmail(email string) (*models.WhitelistUser, error) {
	//data structure to where the data will be written to
	var user models.WhitelistUser

	//choose the database and collection
	collection := repo.Client.Database("podchef").Collection("whitelist")

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

	return &user, nil
}

//GetUserFromWhitelistByID method responsible for getting a user from the database
func (repo *MongoRepository) GetUserFromWhitelistByID(id string) (*models.WhitelistUser, error) {
	//data structure to where the data will be written to
	var user models.WhitelistUser

	//choose the database and collection
	collection := repo.Client.Database("podchef").Collection("whitelist")

	//data to filter the search with
	filter := bson.M{"id": id}

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

	return &user, nil
}

//GetAllUsersFromWhitelist method responsible for all users fro mthe whitelist
func (repo *MongoRepository) GetAllUsersFromWhitelist() ([]models.WhitelistUser, error) {
	//data structure to where the data will be written to
	var users []models.WhitelistUser

	//choose the database and collection
	collection := repo.Client.Database("podchef").Collection("whitelist")

	//data to filter the search with
	filter := bson.M{}

	//call driven adapter responsible for getting a users's data from the database to a cursor
	cur, err := collection.Find(context.Background(), filter)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, &pkg.Error{Err: err, Code: http.StatusNotFound, Message: "No users were found"}
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
func (repo *MongoRepository) InsertUserIntoWhitelist(email string) (*string, error) {
	//data structure containing the data to be inserted
	user := models.WhitelistUser{
		Email: email,
		Data:  time.Now().UTC(),
	}

	//choose the database and collection
	collection := repo.Client.Database("podchef").Collection("whitelist")

	//call driven adapter responsible for inserting a user into the database
	hex, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		//print the error stack
		log.Error(err)

		//return a custom error
		return nil, &pkg.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	id := hex.InsertedID.(primitive.ObjectID).String()

	return &id, nil
}

//DeleteUserFromWhitelistByID method responsible for deleting a user
func (repo *MongoRepository) DeleteUserFromWhitelistByID(id string) (bool, error) {
	//choose the database and collection
	collection := repo.Client.Database("podchef").Collection("whitelist")

	hexID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		//return a custom error
		return false, &pkg.Error{Err: err, Code: http.StatusNotFound, Message: "User not found"}
	}

	//data to filter with
	filter := bson.M{"_id": hexID}

	//call driven adapter responsible for deleting a user from the database
	_, err = collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		//print the error stack
		log.Error(err)

		//return a custom error
		return false, &pkg.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	return true, nil
}
