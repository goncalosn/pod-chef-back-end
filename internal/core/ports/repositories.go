package ports

import (
	appsv1 "k8s.io/api/apps/v1"
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
