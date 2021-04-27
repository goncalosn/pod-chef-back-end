package deployments

import (
	ports "pod-chef-back-end/internal/core/ports"
)

func NewGetService(k8DeploymentsRepository ports.Deployment, k8NamespacesRepository ports.Namespace) *Service {
	return &Service{
		k8DeploymentsRepository: k8DeploymentsRepository,
		k8NamespacesRepository:  k8NamespacesRepository,
	}
}

func (srv *Service) GetDeployments() (interface{}, error) {
	response, err := srv.k8DeploymentsRepository.GetDeployments()

	if err != nil {
		return nil, err
	}

	return response, nil
}
