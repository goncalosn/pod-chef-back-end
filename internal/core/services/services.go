package services

import (
	"pod-chef-back-end/internal/core/ports"
	"pod-chef-back-end/internal/core/services/deployments"
	"pod-chef-back-end/internal/core/services/namespaces"
	"pod-chef-back-end/internal/core/services/nodes"
	"pod-chef-back-end/internal/core/services/pods"
	"pod-chef-back-end/internal/core/services/services"
	"pod-chef-back-end/internal/core/services/users"
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

// ServiceServices stands for kubernetes service
func ServiceServices(kubernetesRepository ports.Service) *services.Service {
	return services.NewService(kubernetesRepository)
}

func VolumeServices(kubernetesRepository ports.Volume) *volumes.Service {
	return volumes.NewService(kubernetesRepository)
}

func UserServices(mongoRepository ports.UserAuth) *users.Service {
	return users.NewService(mongoRepository)
}
