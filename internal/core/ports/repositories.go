package ports

import (
	email "pod-chef-back-end/internal/core/domain/email"
	mongo "pod-chef-back-end/internal/core/domain/mongo"
)

//KubernetesRepository interface holding all the kubernetes respository methods
type KubernetesRepository interface {
	GetNodeByName(name string) (interface{}, error)
	GetNodes() (interface{}, error)

	GetPodsByNodeAndNamespace(node string, namespace string) (interface{}, error)

	CreateDeployment(namespaceUUID string, name string, replicas *int32, image string) (interface{}, error)
	GetDeploymentByNameAndNamespace(name string, namespace string) (interface{}, error)

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
	GetDeploymentByName(name string) (interface{}, error)
	GetAllDeploymentsByUser(userEmail string) (interface{}, error)
	InsertDeployment(name string, namespace string, userEmail string, dockerImage string) (interface{}, error)
	DeleteDeploymentByName(name string) (interface{}, error)
	GetUserByEmail(email string) (*mongo.User, error)
	GetAllUsers() (interface{}, error)
	InsertUser(email string, hash string, name string, role string) (*mongo.User, error)
	DeleteUserByEmail(email string) (interface{}, error)
	GetUserFromWhitelistByEmail(email string) (interface{}, error)
	GetAllUsersFromWhitelist() (interface{}, error)
	InsertUserIntoWhitelist(email string) (interface{}, error)
	DeleteUserFromWhitelistByEmail(email string) (interface{}, error)
}

//EmailRepository interface holding all the email respository methods
type EmailRepository interface {
	SendEmailSMTP(to string, data *email.Email, template string) (bool, error)
}
