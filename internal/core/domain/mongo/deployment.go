package mongo

//Deployment struct with the needed values from the user's deployment
type Deployment struct {
	id          string `bson:",inline"`
	Name        string `bson:"name"`
	Namespace   string `bson:"namespace"`
	UserEmail   string `bson:"user_email"`
	CreatedAt   string `bson:"created_at"`
	DockerImage string `bson:"docker_image"`
}
