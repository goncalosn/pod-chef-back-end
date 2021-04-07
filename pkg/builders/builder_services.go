package pkg

import (
	. "pod-chef-back-end/pkg/kubernetes/nodes"
	. "pod-chef-back-end/pkg/kubernetes/pods"

	"k8s.io/client-go/kubernetes"
)

type ServicesContainer struct {
	NodeService *NodeService
	PodService  *PodService
}

func BuildServices(kubernetesClient *kubernetes.Clientset) *ServicesContainer {
	return &ServicesContainer{
		NodeService: NewNodeService(kubernetesClient),
		PodService:  NewPodService(kubernetesClient),
	}
}
