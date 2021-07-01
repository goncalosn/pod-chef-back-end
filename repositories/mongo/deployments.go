package mongo

import (
	"context"
	"net/http"
	"time"

	models "pod-chef-back-end/internal/core/domain/mongo"
	pkg "pod-chef-back-end/pkg"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//GetDeploymentByUUID method responsible for getting a deployment
func (repo *MongoRepository) GetDeploymentByUUID(uuid string) (*models.Deployment, error) {
	//data structure to where the data will be written to
	var deployment models.Deployment

	//choose the database and collection
	collection := repo.Client.Database("podchef").Collection("deployments")

	//data to filter the search with
	filter := bson.D{{"uuid", uuid}}

	//call driven adapter responsible for getting a user data from the database
	err := collection.FindOne(context.Background(), filter, &options.FindOneOptions{Projection: bson.M{"_id": 0}}).Decode(&deployment)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, &pkg.Error{Err: err, Code: http.StatusNotFound, Message: "Deployment not found"}
		}

		//print the error stack
		log.Error(err)

		//return a custom error
		return nil, &pkg.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}

	}

	return &deployment, nil
}

//GetDeploymentsFromUser method responsible for getting deployment
func (repo *MongoRepository) GetDeploymentsFromUser(id string) ([]models.Deployment, error) {
	//data structure to where the data will be written to
	var deployments []models.Deployment

	//choose the database and collection
	collection := repo.Client.Database("podchef").Collection("deployments")

	//data to filter the search with
	filter := bson.M{"user": id}

	//call driven adapter responsible for getting a deployment's data from the database to a cursor
	cur, err := collection.Find(context.Background(), filter, &options.FindOptions{Projection: bson.M{"_id": 0, "user": 0}})

	// this somehow doesnt work
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, &pkg.Error{Err: err, Code: http.StatusNotFound, Message: "No deployments were found"}
		}

		//print the error stack
		log.Error(err)

		//return a custom error
		return nil, &pkg.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	//decode all the data from the cursor into the deplloyment's structure
	err = cur.All(context.Background(), &deployments)

	if err != nil {
		//print the error stack
		log.Error(err)

		//return a custom error
		return nil, &pkg.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	return deployments, nil
}

//InsertDeployment method responsible for inserting a deployment in the database
func (repo *MongoRepository) InsertDeployment(uuid string, user string, image string) (bool, error) {
	//data structure containing the data to be inserted
	deployment := &models.Deployment{
		UUID:      uuid,
		User:      user,
		CreatedAt: time.Now().UTC().String(),
		Image:     image,
	}

	//choose the database and collection
	collection := repo.Client.Database("podchef").Collection("deployments")

	//call driven adapter responsible for inserting a user into the database
	_, err := collection.InsertOne(context.Background(), deployment)
	if err != nil {
		//print the error stack
		log.Error(err)

		//return a custom error
		return false, &pkg.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	return true, nil
}

//DeleteDeploymentByUUID method responsible for deleting a deployment
func (repo *MongoRepository) DeleteDeploymentByUUID(uuid string) (bool, error) {
	//choose the database and collection
	collection := repo.Client.Database("podchef").Collection("deployments")

	//data to filter with
	filter := bson.D{{"uuid", uuid}}

	//call driven adapter responsible for deleting a deployment from the database
	_, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		//print the error stack
		log.Error(err)

		//return a custom error
		return false, &pkg.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	return true, nil
}
