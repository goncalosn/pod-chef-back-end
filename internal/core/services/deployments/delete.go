package deployments

import (
	ports "pod-chef-back-end/internal/core/ports"
)

func NewDeleteService(k8DeploymentsRepository ports.Deployment, k8NamespacesRepository ports.Namespace) *Service {
	return &Service{
		k8DeploymentsRepository: k8DeploymentsRepository,
		k8NamespacesRepository:  k8NamespacesRepository,
	}
}

func (srv *Service) DeleteDeployment(name string) (interface{}, error) {
	response, err := srv.k8DeploymentsRepository.DeleteDeployment(name)

	if err != nil {
		return nil, err
	}

	return response, nil
}
