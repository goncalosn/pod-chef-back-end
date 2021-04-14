package deployments

import (
	ports "pod-chef-back-end/internal/core/ports"
)

func NewDeleteService(kubernetesRepository ports.Deployment) *Service {
	return &Service{
		kubernetesRepository: kubernetesRepository,
	}
}

func (srv *Service) DeleteDeployment(name string) (interface{}, error) {
	response, err := srv.kubernetesRepository.DeleteDeployment(name)

	if err != nil {
		return nil, err
	}

	return response, nil
}
