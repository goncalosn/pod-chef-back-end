package mongodb

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Client() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()
	// os.Setenv("MONGO_URI", "mongodb+srv://admin:yYJFMpHCrz5uj327@cluster0.oy5rg.mongodb.net/main?retryWrites=true&w=majority")
	uri := os.Getenv("MONGO_URI")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	return client
}