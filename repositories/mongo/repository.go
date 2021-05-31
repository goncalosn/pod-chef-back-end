package services

import (
	"context"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
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
	return client
}
