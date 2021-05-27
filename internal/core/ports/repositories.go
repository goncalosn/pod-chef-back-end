package ports

type Node interface {
	GetNode(name string) (interface{}, error)
	GetNodes() (interface{}, error)
}

type Pod interface {
	GetPodsByNodeAndNamespace(node string, namespace string) (interface{}, error)
}

type Deployment interface {
	CreateDeployment(namespaceUuid string, name string, replicas *int32, image string) (interface{}, error)
	GetDeployments() (interface{}, error)
	DeleteDeployment(name string) (interface{}, error)
}

type Namespace interface {
	GetNamespaces() ([]string, error)
	CreateNamespace(name string) (interface{}, error)
	DeleteNamespace(name string) (interface{}, error)
}

// Service stands for kubernetes service
type Service interface {
	GetServicesByNamespace(namespace string) (interface{}, error)
	GetServiceByNameAndNamespace(name string, namespace string) (interface{}, error)
	CreateClusterIPService(namespace string, name string) (interface{}, error)
}

type Ingress interface {
	GetIngress(name string) (interface{}, error)
	CreateIngress(namespace string, name string, host string) (interface{}, error)
}
