package deployments

import (
	ports "pod-chef-back-end/internal/core/ports"
)

func NewGetService(kubernetesRepository ports.Deployment) *Service {
	return &Service{
		kubernetesRepository: kubernetesRepository,
	}
}

func (srv *Service) GetDeployments() (interface{}, error) {
	response, err := srv.kubernetesRepository.GetDeployments()

	if err != nil {
		return nil, err
	}

	return response, nil
}
