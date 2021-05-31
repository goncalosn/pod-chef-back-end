package services

import (
	"context"
	"net/http"
	"time"

	models "pod-chef-back-end/internal/core/domain/mongo"
	pkg "pod-chef-back-end/pkg"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//GetDeploymentByName method responsible for getting a deployment
func (repo *MongoRepository) GetDeploymentByName(name string) (interface{}, error) {
	//data structure to where the data will be written to
	var deployment *models.Deployment

	//choose the database and collection
	collection := repo.Client.Database("podchef").Collection("deployments")

	//data to filter the search with
	filter := bson.D{{"name", name}}

	//call driven adapter responsible for getting a user data from the database
	err := collection.FindOne(context.Background(), filter).Decode(&deployment)
	if err != nil && err != mongo.ErrNoDocuments {
		//print the error stack
		log.Error(err)

		//return a custom error
		return nil, &pkg.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	return deployment, nil
}

//GetAllDeploymentsByUser method responsible for getting deployment
func (repo *MongoRepository) GetAllDeploymentsByUser(userEmail string) (interface{}, error) {
	//data structure to where the data will be written to
	var deployments *[]models.Deployment

	//choose the database and collection
	collection := repo.Client.Database("podchef").Collection("deployments")

	//data to filter the search with
	filter := bson.D{{"user_email", userEmail}}

	//call driven adapter responsible for getting a deployment's data from the database to a cursor
	cur, err := collection.Find(context.Background(), filter)

	if err != nil {
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
func (repo *MongoRepository) InsertDeployment(name string, namespace string, userEmail string, dockerImage string) (interface{}, error) {
	//data structure containing the data to be inserted
	deployment := &models.Deployment{
		Name:        name,
		Namespace:   namespace,
		UserEmail:   userEmail,
		CreatedAt:   time.Now().UTC().String(),
		DockerImage: dockerImage,
	}

	//choose the database and collection
	collection := repo.Client.Database("podchef").Collection("deployments")

	//call driven adapter responsible for inserting a user into the database
	response, err := collection.InsertOne(context.Background(), deployment)
	if err != nil {
		//print the error stack
		log.Error(err)

		//return a custom error
		return nil, &pkg.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	return response, nil
}

//DeleteDeploymentByName method responsible for deleting a deployment
func (repo *MongoRepository) DeleteDeploymentByName(name string) (interface{}, error) {
	//choose the database and collection
	collection := repo.Client.Database("podchef").Collection("deployments")

	//data to filter with
	filter := bson.D{{"name", name}}

	//call driven adapter responsible for deleting a deployment from the database
	_, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		//print the error stack
		log.Error(err)

		//return a custom error
		return nil, &pkg.Error{Err: err, Code: http.StatusInternalServerError, Message: "Internal error"}
	}

	return nil, nil
}
