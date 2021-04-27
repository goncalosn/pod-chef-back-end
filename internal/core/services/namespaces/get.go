package namespaces

import (
	ports "pod-chef-back-end/internal/core/ports"
)

func NewGetService(kubernetesRepository ports.Namespace) *Service {
	return &Service{
		kubernetesRepository: kubernetesRepository,
	}
}

func (srv *Service) GetNamespaces() (interface{}, error) {
	response, err := srv.kubernetesRepository.GetNamespaces()

	if err != nil {
		return nil, err
	}

	return response, nil
}
