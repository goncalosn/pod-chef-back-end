package ingresses

import "k8s.io/client-go/kubernetes"

//this service's dependencies
type KubernetesClient struct {
	Clientset *kubernetes.Clientset
}

func New(clientset *kubernetes.Clientset) *KubernetesClient {
	return &KubernetesClient{
		Clientset: clientset,
	}
}
