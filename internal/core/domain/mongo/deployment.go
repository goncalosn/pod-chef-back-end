package mongo

//Deployment struct with the needed values from the user's deployment
type Deployment struct {
	id        string `bson:",inline"`
	UUID      string `bson:"uuid"`
	User      string `bson:"user"`
	CreatedAt string `bson:"created_at"`
	Image     string `bson:"image"`
}
