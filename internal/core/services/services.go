package services

import (
	ports "pod-chef-back-end/internal/core/ports"
	nodes "pod-chef-back-end/internal/core/services/nodes"
	pods "pod-chef-back-end/internal/core/services/pods"
)

func NewNodeServices(kubernetesRepository ports.Node) {
	nodes.NewGetService(kubernetesRepository)
}

func NewPodServices(kubernetesRepository ports.Pod) {
	pods.NewGetService(kubernetesRepository)
}
