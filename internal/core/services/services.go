package services

import (
	ports "pod-chef-back-end/internal/core/ports"
	deployments "pod-chef-back-end/internal/core/services/deployments"
	nodes "pod-chef-back-end/internal/core/services/nodes"
	pods "pod-chef-back-end/internal/core/services/pods"
	services "pod-chef-back-end/internal/core/services/services"
)

func NodeServices(kubernetesRepository ports.Node) *nodes.Service {
	return nodes.NewService(kubernetesRepository)
}

func PodServices(kubernetesRepository ports.Pod) *pods.Service {
	return pods.NewService(kubernetesRepository)
}

func DeploymentServices(k8DeploymentRepository ports.Deployment, k8ServicesRepository ports.Service) *deployments.Service {
	return deployments.NewService(k8DeploymentRepository, k8ServicesRepository)
}

// Service stands for kubernetes service
func ServiceServices(kubernetesRepository ports.Service) *services.Service {
	return services.NewService(kubernetesRepository)
}
