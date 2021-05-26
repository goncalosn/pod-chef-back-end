package deployments

import "pod-chef-back-end/internal/core/ports"

type Service struct {
	k8DeploymentsRepository ports.Deployment
	k8NamespacesRepository  ports.Namespace
	k8ServicesRepository    ports.Service
	k8IngressesRepository   ports.Ingress
}

func NewService(k8DeploymentsRepository ports.Deployment, k8NamespacesRepository ports.Namespace, k8ServicesRepository ports.Service,
	k8IngressesRepository ports.Ingress) *Service {
	return &Service{
		k8DeploymentsRepository: k8DeploymentsRepository,
		k8NamespacesRepository:  k8NamespacesRepository,
		k8ServicesRepository:    k8ServicesRepository,
		k8IngressesRepository:   k8IngressesRepository,
	}
}
