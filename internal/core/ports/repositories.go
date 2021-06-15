package ports

import (
	email "pod-chef-back-end/internal/core/domain/email"
	models "pod-chef-back-end/internal/core/domain/mongo"
	mongo "pod-chef-back-end/internal/core/domain/mongo"
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
	GetAllUsers() (*[]models.User, error)
	InsertUser(email string, hash string, name string, role string) (*mongo.User, error)
	DeleteUserByEmail(email string) (bool, error)
	UpdateUserPassword(email string, hash string) (bool, error)
	UpdateUserRole(email string, role string) (bool, error) 
	UpdateUserName(email string, name string) (bool, error)

	GetUserFromWhitelistByEmail(email string) (interface{}, error)
	GetAllUsersFromWhitelist() ([]models.User, error)
	InsertUserIntoWhitelist(email string) (bool, error)
	DeleteUserFromWhitelistByEmail(email string) (bool, error)

	GetDeploymentByUUID(uuid string) (*models.Deployment, error)
	GetDeploymentsFromUser(email string) ([]models.Deployment, error)
	InsertDeployment(uuid string, email string, image string) (bool, error)
	DeleteDeploymentByUUID(uuid string) (bool, error)
}

//EmailRepository interface holding all the email respository methods
type EmailRepository interface {
	SendEmailSMTP(to string, data *email.Email, template string) (bool, error)
}
