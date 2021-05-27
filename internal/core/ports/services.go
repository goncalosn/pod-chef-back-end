package ports

type NodeServices interface {
	GetNode(name string) (interface{}, error)
	GetNodes() (interface{}, error)
}

type PodServices interface {
	GetPodsByNodeAndNamespace(node string, namespace string) (interface{}, error)
}

type DeploymentServices interface {
	CreateDeployment(token string, replicas *int32, image string) (interface{}, error)
	GetDeployments() (interface{}, error)
	DeleteDeployment(name string) (interface{}, error)
}

type NamespaceServices interface {
	GetNamespaces() (interface{}, error)
	CreateNamespace(name string) (interface{}, error)
	DeleteNamespace(name string) (interface{}, error)
}

// Service stands for kubernetes service
type ServiceServices interface {
	GetServicesByNamespace(namespace string) (interface{}, error)
	GetServiceByNameAndNamespace(name string, namespace string) (interface{}, error)
	CreateClusterIPService(namespace string, name string) (interface{}, error)
}

type IngressServices interface {
	GetIngress(name string) (interface{}, error)
	CreateIngress(namespace string, name string, host string) (interface{}, error)
}
