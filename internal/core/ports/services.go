package ports

import (
	"mime/multipart"
)

type NodeServices interface {
	GetNode(name string) (interface{}, error)
	GetNodes() (interface{}, error)
}

type PodServices interface {
	GetPodsByNodeAndNamespace(node string, namespace string) (interface{}, error)
}

type DeploymentServices interface {
	CreateDefaultDeployment(name string, replicas *int32, image string) (interface{}, error)
	CreateFileDeployment(*multipart.FileHeader) (interface{}, error)
	GetDeployments() (interface{}, error)
	DeleteDeployment(name string) (interface{}, error)
}

type NamespaceServices interface {
	GetNamespaces() (interface{}, error)
}

// ServiceServices stands for kubernetes service
type ServiceServices interface {
	GetServicesByNamespace(namespace string) (interface{}, error)
	GetServiceByNameAndNamespace(name string, namespace string) (interface{}, error)
}

type VolumeServices interface {
	GetVolumes() (interface{}, error)
}

type UserServices interface {
	Register(username string, email string, password string) (interface{}, error)
	Authenticate(email string, password string) (interface{}, error)
}
