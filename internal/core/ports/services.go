package ports

type KubernetesServices interface {
	GetNode(name string) (interface{}, error)
	GetNodes() (interface{}, error)

	GetPodsByNodeAndNamespace(node string, namespace string) (interface{}, error)

	CreateDeployment(token string, replicas *int32, image string) (interface{}, error)
	GetDeployments() (interface{}, error)
	DeleteDeployment(name string) (interface{}, error)

	GetNamespaces() (interface{}, error)
	CreateNamespace(name string) (interface{}, error)
	DeleteNamespace(name string) (interface{}, error)

	GetServicesByNamespace(namespace string) (interface{}, error)
	GetServiceByNameAndNamespace(name string, namespace string) (interface{}, error)
	CreateClusterIPService(namespace string, name string) (interface{}, error)

	GetIngress(name string) (interface{}, error)
	CreateIngress(namespace string, name string, host string) (interface{}, error)
}
