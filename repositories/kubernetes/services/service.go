package services

import "k8s.io/client-go/kubernetes"

//this service's dependencies
type KubernetesClient struct {
	Clientset *kubernetes.Clientset
}

//service in charge of dealing with GET requests and nodes
func New(clientset *kubernetes.Clientset) *KubernetesClient {
	return &KubernetesClient{
		Clientset: clientset,
	}
}
