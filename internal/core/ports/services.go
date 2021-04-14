package ports

type NodeServices interface {
	GetNode(name string) (interface{}, error)
	GetNodes() (interface{}, error)
}

type PodServices interface {
	GetPodsByNodeAndNamespace(node string, namespace string) (interface{}, error)
}
