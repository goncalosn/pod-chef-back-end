package ports

//KubernetesRepository interface holding all the kubernetes respository methods
type KubernetesRepository interface {
	GetNodeByName(name string) (interface{}, error)
	GetNodes() (interface{}, error)

	GetPodsByNodeAndNamespace(node string, namespace string) (interface{}, error)

	CreateDeployment(namespaceUUID string, name string, replicas *int32, image string) (interface{}, error)
	GetDeploymentByNameAndNamespace(name string, namespace string) (interface{}, error)

	GetNamespaces() ([]string, error)
	CreateNamespace(name string) (interface{}, error)
	DeleteNamespace(name string) (interface{}, error)

	GetServicesByNamespace(namespace string) (interface{}, error)
	GetServiceByNameAndNamespace(name string, namespace string) (interface{}, error)
	CreateClusterIPService(namespace string, name string) (interface{}, error)

	GetIngressByNameAndNamespace(name string, namespace string) (interface{}, error)
	CreateIngress(namespace string, name string, host string) (interface{}, error)
}
