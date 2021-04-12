package ports

type Node interface {
	GetNode(name string) (interface{}, error)
	GetNodes() (interface{}, error)
}

type Pod interface {
	GetPodsByNodeAndNamespace(node string, namespace string) (interface{}, error)
}
