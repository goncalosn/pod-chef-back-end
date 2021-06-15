package ports

import (
	models "pod-chef-back-end/internal/core/domain/mongo"
)

//KubernetesServices interface holding all the kubernetes services
type KubernetesServices interface {
	GetNodeByName(name string) (interface{}, error)
	GetNodes() (interface{}, error)

	GetPodsByNodeAndNamespace(node string, namespace string) (interface{}, error)

	CreateDeployment(email string, replicas *int32, image string) (interface{}, error)
	GetDeploymentsByUser(email string) (interface{}, error)
	GetDeploymentByUserAndName(email string, name string) (interface{}, error)
	DeleteDeploymentByUserAndUUID(email string, uuid string) (interface{}, error)

	CreateNamespace(name string) (interface{}, error)
	DeleteNamespace(name string) (interface{}, error)

	GetServicesByNamespace(namespace string) (interface{}, error)
	GetServiceByNameAndNamespace(name string, namespace string) (interface{}, error)
	CreateClusterIPService(namespace string, name string) (interface{}, error)

	GetIngressByName(name string, namespace string) (interface{}, error)
	CreateIngress(namespace string, name string, host string) (interface{}, error)
}

//MongoServices interface holding all the mongo services
type MongoServices interface {
	GetUserByEmail(email string) (*models.User, error)
	GetAllUsers() (*[]models.User, error)
	InsertUser(email string, password string, name string) (*models.User, error)
	DeleteUser(name string) (bool, error)
	UpdateSelfPassword(email string, hash string) (bool, error)
	ResetUserPassword(to string, password string) (bool, error)
	UpdateUserRole(email string, role string) (bool, error)
	UpdateUserName(email string, name string) (bool, error)

	GetAllUsersFromWhitelist() (*[]models.User, error)
	InviteUserToWhitelist(to string) (bool, error)
	RemoveUserFromWhitelist(to string) (bool, error)
}
