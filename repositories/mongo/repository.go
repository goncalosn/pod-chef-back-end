package mongo

import (
	"context"
	models "pod-chef-back-end/internal/core/domain/mongo"
	pkg "pod-chef-back-end/pkg"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//MongoRepository mongo client responsible for acessing the database
type MongoRepository struct {
	Client *mongo.Client
}

//NewMongoRepository new connection to the mongo database
func NewMongoRepository(viper *viper.Viper) *MongoRepository {
	return &MongoRepository{
		Client: Client(viper),
	}
}

//Client responsible for creating the connection to the mongo database
func Client(viper *viper.Viper) *mongo.Client {
	username := viper.Get("DB_USER").(string)
	password := viper.Get("DB_PASSWORD").(string)

	defaultAdminEmail := viper.Get("DEFAULT_ADMIN_EMAIL").(string)
	defaultAdminPassword := viper.Get("DEFAULT_ADMIN_PASSWORD").(string)

	log.Info("creating conection to the database")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Info("connecting to the database")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://"+username+":"+password+"@cluster0.tsevj.mongodb.net/development?w=majority",
	))
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	log.Info("connection to the database sucessful")

	createDefaultAdminAccount(client, defaultAdminEmail, defaultAdminPassword)

	return client
}

func createDefaultAdminAccount(client *mongo.Client, adminEmail string, adminPass string) {
	collection := client.Database("podchef").Collection("users")
	// check if exists users in collection
	res, err := collection.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		if err != mongo.ErrNoDocuments {
			log.Fatal(err)
		}
	}

	if res == 0 {
		crypt := pkg.EncryptPassword(adminPass)
		user := &models.User{
			Email: adminEmail,
			Hash:  string(crypt),
			Name:  adminEmail,
			Role:  "admin",
			Date:  time.Now().UTC(),
		}
		_, err := collection.InsertOne(context.Background(), user)

		if err != nil {
			//type assertion of custom Error to default error
			mongoError := err.(*pkg.Error)
			//return the error sent by the service
			log.Fatal(mongoError)
		}
		log.Info("default admin inserted")
	}
}
