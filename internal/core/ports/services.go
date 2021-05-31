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
	InsertUser(email string, password string, name string, role string) (*models.User, error)
	DeleteUser(name string) (interface{}, error)
}
