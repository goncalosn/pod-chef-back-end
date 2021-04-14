package kubernetes

import (
	deployments "pod-chef-back-end/repositories/kubernetes/deployments"
	nodes "pod-chef-back-end/repositories/kubernetes/nodes"
	pods "pod-chef-back-end/repositories/kubernetes/pods"
)

type kubernetesRepository struct {
	Pods        *pods.KubernetesClient
	Nodes       *nodes.KubernetesClient
	Deployments *deployments.KubernetesClient
}

func KubernetesRepository() *kubernetesRepository {
	client := Client()

	return &kubernetesRepository{
		Pods:        pods.New(client),
		Nodes:       nodes.New(client),
		Deployments: deployments.New(client),
	}
}
