package ports

import "mime/multipart"

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
	CheckRepeatedDeployName(name string) (bool, error)
	DeleteDeployment(name string) (interface{}, error)
}
