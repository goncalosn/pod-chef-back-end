package services

import (
	ports "pod-chef-back-end/internal/core/ports"
	nodes "pod-chef-back-end/internal/core/services/nodes"
	pods "pod-chef-back-end/internal/core/services/pods"
)

func NodeServices(kubernetesRepository ports.Node) *nodes.Service {
	return nodes.NewGetService(kubernetesRepository)
}

func PodServices(kubernetesRepository ports.Pod) *pods.Service {
	return pods.NewGetService(kubernetesRepository)
}
