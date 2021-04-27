package kubernetes

import (
	deployments "pod-chef-back-end/repositories/kubernetes/deployments"
	namespaces "pod-chef-back-end/repositories/kubernetes/namespaces"
	nodes "pod-chef-back-end/repositories/kubernetes/nodes"
	pods "pod-chef-back-end/repositories/kubernetes/pods"
	services "pod-chef-back-end/repositories/kubernetes/services"
	volumes "pod-chef-back-end/repositories/kubernetes/volumes"
)

type kubernetesRepository struct {
	Pods        *pods.KubernetesClient
	Nodes       *nodes.KubernetesClient
	Deployments *deployments.KubernetesClient
	Namespaces  *namespaces.KubernetesClient
	Services    *services.KubernetesClient
	Volumes     *volumes.KubernetesClient
}

func KubernetesRepository() *kubernetesRepository {
	client := Client()

	return &kubernetesRepository{
		Pods:        pods.New(client),
		Nodes:       nodes.New(client),
		Deployments: deployments.New(client),
		Namespaces:  namespaces.New(client),
		Services:    services.New(client),
		Volumes:     volumes.New(client),
	}
}
