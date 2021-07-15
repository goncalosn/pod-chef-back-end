package ports

import (
	models "pod-chef-back-end/internal/core/domain/mongo"
	mongo "pod-chef-back-end/internal/core/domain/mongo"
)

//KubernetesRepository interface holding all the kubernetes respository methods
type KubernetesRepository interface {
	GetNodeByName(name string) (interface{}, error)
	GetNodes() (interface{}, error)

	GetPodsByNodeAndNamespace(node string, namespace string) (interface{}, error)

	GetDeploymentByNameAndNamespace(name string, namespace string) (interface{}, error)
	CreateDeployment(namespace string, name string, replicas *int32, image string, containerPort int32) (interface{}, error)

	CreateNamespace(name string) (interface{}, error)
	DeleteNamespace(name string) (interface{}, error)

	GetServicesByNamespace(namespace string) (interface{}, error)
	GetServiceByNameAndNamespace(name string, namespace string) (interface{}, error)
	CreateClusterIPService(namespace string, name string, containerPort int32) (interface{}, error)

	GetIngressByNameAndNamespace(name string, namespace string) (interface{}, error)
	CreateIngress(namespace string, name string, uuid string) (interface{}, error)
}

//MongoRepository interface holding all the mongo respository methods
type MongoRepository interface {
	GetUserByEmail(email string) (*mongo.User, error)
	GetUserByID(id string) (*mongo.User, error)
	GetAllUsers() (*[]models.User, error)
	InsertUser(email string, hash string, name string, role string) (*mongo.User, error)
	DeleteUserByID(id string) (bool, error)
	UpdateUserPassword(id string, hash string) (bool, error)
	UpdateUserRole(id string, role string) (bool, error)
	UpdateUserName(id string, name string) (bool, error)

	GetUserFromWhitelistByEmail(email string) (*models.WhitelistUser, error)
	GetUserFromWhitelistByID(id string) (*models.WhitelistUser, error)
	GetAllUsersFromWhitelist() ([]models.WhitelistUser, error)
	InsertUserIntoWhitelist(email string) (*string, error)
	DeleteUserFromWhitelistByID(id string) (bool, error)
	DeleteUserFromWhitelistByEmail(email string) (bool, error)

	GetAllDeployments() ([]models.Deployment, error)
	GetDeploymentByUUID(uuid string) (*models.Deployment, error)
	GetDeploymentsFromUser(id string) ([]models.Deployment, error)
	InsertDeployment(uuid string, email string, image string) (bool, error)
	DeleteDeploymentByUUID(uuid string) (bool, error)
}

//EmailRepository interface holding all the email respository methods
type EmailRepository interface {
	SendEmailSMTP(to string, subject string, emailBody string) error
}
