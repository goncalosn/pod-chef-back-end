package ports

//KubernetesServices interface holding all the kubernetes services
type KubernetesServices interface {
	GetNodeByName(name string) (interface{}, error)
	GetNodes() (interface{}, error)

	GetPodsByNodeAndNamespace(node string, namespace string) (interface{}, error)

	CreateDeployment(token string, replicas *int32, image string) (interface{}, error)
	GetDeploymentsByUser(token string) (interface{}, error)
	GetDeploymentByUserAndName(token string, name string) (interface{}, error)

	GetNamespaces() (interface{}, error)
	CreateNamespace(name string) (interface{}, error)
	DeleteNamespace(name string) (interface{}, error)

	GetServicesByNamespace(namespace string) (interface{}, error)
	GetServiceByNameAndNamespace(name string, namespace string) (interface{}, error)
	CreateClusterIPService(namespace string, name string) (interface{}, error)

	GetIngressByName(name string, namespace string) (interface{}, error)
	CreateIngress(namespace string, name string, host string) (interface{}, error)
}
