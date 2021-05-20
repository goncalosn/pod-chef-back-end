package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Client() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		"mongodb+srv://admin:niggapassword@cluster0.oy5rg.mongodb.net/myFirstDatabase?retryWrites=true&w=majority",
	))
	if err != nil {
		log.Fatal(err)
		panic("mongo error")
		//return nil
	}

	return client
}
