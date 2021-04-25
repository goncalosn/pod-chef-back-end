package deployments

import "pod-chef-back-end/internal/core/ports"

type Service struct {
	k8DeploymentsRepository ports.Deployment
	k8NamespacesRepository  ports.Namespace
}

func NewService(k8DeploymentsRepository ports.Deployment, k8NamespacesRepository ports.Namespace) *Service {
	return &Service{
		k8DeploymentsRepository: k8DeploymentsRepository,
		k8NamespacesRepository:  k8NamespacesRepository,
	}
}
