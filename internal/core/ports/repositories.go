package ports

import (
	"pod-chef-back-end/pkg/auth"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
)

type Node interface {
	GetNode(name string) (interface{}, error)
	GetNodes() (interface{}, error)
}

type Pod interface {
	GetPodsByNodeAndNamespace(node string, namespace string) (interface{}, error)
}

type Deployment interface {
	CreateDefaultDeployment(name string, replicas *int32, image string) (interface{}, error)
	CreateFileDeployment(dep *appsv1.Deployment) (interface{}, error)
	GetDeployments() (interface{}, error)
	CheckRepeatedDeployName(name string, namespace string) (bool, error)
	DeleteDeployment(name string) (interface{}, error)
}

type Namespace interface {
	GetNamespaces() ([]string, error)
}

// Service stands for kubernetes service
type Service interface {
	GetServicesByNamespace(namespace string) (interface{}, error)
	GetServiceByNameAndNamespace(name string, namespace string) (interface{}, error)
	CreateService(serv *v1.Service) (interface{}, error)
}

type Volume interface {
	GetVolumes() (interface{}, error)
}

type User interface {
	SignIn(user auth.User) (interface{}, error)
	Login(email string) (interface{}, error)
}
