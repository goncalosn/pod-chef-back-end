package ports

import (
	models "pod-chef-back-end/internal/core/domain/mongo"
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

	GetUserByEmail(email string) (*models.User, error)
	InsertUser(email string, hash [32]byte, name string, role string) (*models.User, error)
	DeleteUserByEmail(email string) (interface{}, error)
}

type UserAuth interface {
	Register(username string, email string, password string) (interface{}, error)
	Authenticate(email string, password string) (string, error)
}
