package deployments

import "k8s.io/client-go/kubernetes"

// KubernetesClient service's dependencies
type KubernetesClient struct {
	Clientset *kubernetes.Clientset
}

// New service in charge of dealing with GET requests and nodes
func New(clientset *kubernetes.Clientset) *KubernetesClient {
	return &KubernetesClient{
		Clientset: clientset,
	}
}
