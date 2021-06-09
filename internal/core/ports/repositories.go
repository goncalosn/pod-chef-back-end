package ports

import (
	email "pod-chef-back-end/internal/core/domain/email"
	models "pod-chef-back-end/internal/core/domain/mongo"
	mongo "pod-chef-back-end/internal/core/domain/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

//KubernetesRepository interface holding all the kubernetes respository methods
type KubernetesRepository interface {
	GetNodeByName(name string) (interface{}, error)
	GetNodes() (interface{}, error)

	GetPodsByNodeAndNamespace(node string, namespace string) (interface{}, error)

	GetDeploymentByNameAndNamespace(name string, namespace string) (interface{}, error)
	CreateDeployment(namespace string, name string, replicas *int32, image string) (interface{}, error)

	CreateNamespace(name string) (interface{}, error)
	DeleteNamespace(name string) (interface{}, error)

	GetServicesByNamespace(namespace string) (interface{}, error)
	GetServiceByNameAndNamespace(name string, namespace string) (interface{}, error)
	CreateClusterIPService(namespace string, name string) (interface{}, error)

	GetIngressByNameAndNamespace(name string, namespace string) (interface{}, error)
	CreateIngress(namespace string, name string, host string) (interface{}, error)
}

//MongoRepository interface holding all the mongo respository methods
type MongoRepository interface {
	GetUserByEmail(email string) (*mongo.User, error)
	GetAllUsers() (interface{}, error)
	InsertUser(email string, hash string, name string, role string) (*mongo.User, error)
	DeleteUserByEmail(email string) (interface{}, error)

	GetUserFromWhitelistByEmail(email string) (interface{}, error)
	GetAllUsersFromWhitelist() (interface{}, error)
	InsertUserIntoWhitelist(email string) (interface{}, error)
	DeleteUserFromWhitelistByEmail(email string) (interface{}, error)

	GetDeploymentByUUID(uuid string) (*models.Deployment, error)
	GetAllDeploymentsByUser(user string) (*[]bson.M, error)
	InsertDeployment(uuid string, user string, image string) (interface{}, error)
	DeleteDeploymentByUUID(uuid string) (bool, error)
}

//EmailRepository interface holding all the email respository methods
type EmailRepository interface {
	SendEmailSMTP(to string, data *email.Email, template string) (bool, error)
}
