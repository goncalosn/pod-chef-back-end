package mongo

import "time"

//Deployment struct with the needed values from the user's deployment
type Deployment struct {
	id        string    `bson:"_id" json:"_id"`
	UUID      string    `bson:"uuid" json:"uuid"`
	User      string    `bson:"user" json:"user"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	Image     string    `bson:"image" json:"image"`
}
