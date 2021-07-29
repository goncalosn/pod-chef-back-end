package mongo

import (
	"context"
	"net/http"
	"time"

	models "pod-chef-back-end/internal/core/domain/mongo"
	"pod-chef-back-end/pkg"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

//GetUserByID method responsible for getting a user from the database
func (repo *MongoRepository) GetUserByID(id string) (*models.User, error) {
	//data structure to where the data will be written to
	var user *models.User

	//choose the database and collection
	collection := repo.Client.Database("podchef").Collection("users")
	hexID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		//return a custom error
		return nil, &pkg.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	//data to filter the search with
	filter := bson.M{"_id": hexID}

	//call driven adapter responsible for getting a user data from the database
	err = collection.FindOne(context.Background(), filter).Decode(&user)

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
func (repo *MongoRepository) GetAllUsers() (*[]models.User, error) {
	//data structure to where the data will be written to
	var users []models.User

	//choose the database and collection
	collection := repo.Client.Database("podchef").Collection("users")

	//data to filter the search with
	filter := bson.D{}

	//call driven adapter responsible for getting a users's data from the database to a cursor
	cur, err := collection.Find(context.Background(), filter, &options.FindOptions{Projection: bson.M{"hash": 0}})

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &users, nil
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

	return &users, nil
}

//InsertUser method responsible for inserting a user in the database
func (repo *MongoRepository) InsertUser(email string, hash string, name string, role string) (*models.User, error) {
	//data structure containing the data to be inserted
	user := &models.User{
		Email: email,
		Hash:  hash,
		Name:  name,
		Role:  role,
		Date:  time.Now().UTC(),
	}

	//choose the database and collection
	collection := repo.Client.Database("podchef").Collection("users")

	//insert user with bson because of automatic id generator
	hex, err := collection.InsertOne(context.Background(), bson.M{"email": email, "hash": hash, "name": name, "role": role, "date": time.Now().UTC()})

	if err != nil {
		//print the error stack
		log.Error(err)

		//return a custom error
		return nil, &pkg.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	user.ID = hex.InsertedID.(primitive.ObjectID).Hex()
	return user, nil
}

//DeleteUserByID method responsible for deleting a user
func (repo *MongoRepository) DeleteUserByID(id string) (bool, error) {
	//choose the database and collection
	collection := repo.Client.Database("podchef").Collection("users")

	hexID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		//return a custom error
		return false, &pkg.Error{Err: err, Code: http.StatusNotFound, Message: "User not found"}
	}

	res, err := repo.DeleteAllDeploymentFromUser(id)
	if err != nil {
		// if res is true it means that the user never deployed an app,
		// so, it successfully searched
		if res {
			return true, nil
		}
		// if res is false than there is a critical error
		log.Error(err)
		return false, &pkg.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
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

//UpdateUserPassword method responsible for updating a user password
func (repo *MongoRepository) UpdateUserPassword(id string, hash string) (bool, error) {
	//choose the database and collection
	collection := repo.Client.Database("podchef").Collection("users")

	hexID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		//return a custom error
		return false, &pkg.Error{Err: err, Code: http.StatusNotFound, Message: "User not found"}
	}

	//data to filter with
	filter := bson.M{"_id": hexID}
	//data to update
	update := bson.D{
		{"$set", bson.D{{"hash", hash}}},
	}

	//call driven adapter responsible for updating a fild  from the database
	_, err = collection.UpdateOne(context.TODO(), filter, update)

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

//UpdateUserRole method responsible for updating a user role
func (repo *MongoRepository) UpdateUserRole(id string, role string) (bool, error) {
	//choose the database and collection
	collection := repo.Client.Database("podchef").Collection("users")

	hexID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		//return a custom error
		return false, &pkg.Error{Err: err, Code: http.StatusNotFound, Message: "User not found"}
	}

	//data to filter with
	filter := bson.M{"_id": hexID}
	//data to update
	update := bson.D{
		{"$set", bson.D{{"role", role}}},
	}

	//call driven adapter responsible for updating a fild  from the database
	_, err = collection.UpdateOne(context.TODO(), filter, update)

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

//UpdateUserName method responsible for updating a user role
func (repo *MongoRepository) UpdateUserName(id string, name string) (bool, error) {
	//choose the database and collection
	collection := repo.Client.Database("podchef").Collection("users")

	hexID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		//return a custom error
		return false, &pkg.Error{Err: err, Code: http.StatusNotFound, Message: "User not found"}
	}

	//data to filter with
	filter := bson.M{"_id": hexID}

	//data to update
	update := bson.D{
		{"$set", bson.D{{"name", name}}},
	}

	//call driven adapter responsible for updating a fild  from the database
	_, err = collection.UpdateOne(context.TODO(), filter, update)

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
