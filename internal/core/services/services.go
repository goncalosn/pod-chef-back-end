package services

import (
	ports "pod-chef-back-end/internal/core/ports"
	deployments "pod-chef-back-end/internal/core/services/deployments"
	namespaces "pod-chef-back-end/internal/core/services/namespaces"
	nodes "pod-chef-back-end/internal/core/services/nodes"
	pods "pod-chef-back-end/internal/core/services/pods"
	services "pod-chef-back-end/internal/core/services/services"
	volumes "pod-chef-back-end/internal/core/services/volumes"
)

func NodeServices(kubernetesRepository ports.Node) *nodes.Service {
	return nodes.NewService(kubernetesRepository)
}

func PodServices(kubernetesRepository ports.Pod) *pods.Service {
	return pods.NewService(kubernetesRepository)
}

func DeploymentServices(k8DeploymentRepository ports.Deployment, k8NamespacesRepository ports.Namespace, k8ServicesRepository ports.Service) *deployments.Service {
	return deployments.NewService(k8DeploymentRepository, k8NamespacesRepository, k8ServicesRepository)
}

func NamespaceServices(kubernetesRepository ports.Namespace) *namespaces.Service {
	return namespaces.NewService(kubernetesRepository)
}

// Service stands for kubernetes service
func ServiceServices(kubernetesRepository ports.Service) *services.Service {
	return services.NewService(kubernetesRepository)
}

func VolumeServices(kubernetesRepository ports.Volume) *volumes.Service {
	return volumes.NewService(kubernetesRepository)
}
