package services

import (
	ports "pod-chef-back-end/internal/core/ports"
	deployments "pod-chef-back-end/internal/core/services/deployments"
	nodes "pod-chef-back-end/internal/core/services/nodes"
	pods "pod-chef-back-end/internal/core/services/pods"
)

func NodeServices(kubernetesRepository ports.Node) *nodes.Service {
	return nodes.NewService(kubernetesRepository)
}

func PodServices(kubernetesRepository ports.Pod) *pods.Service {
	return pods.NewService(kubernetesRepository)
}

func DeploymentServices(kubernetesRepository ports.Deployment) *deployments.Service {
	return deployments.NewService(kubernetesRepository)
}
