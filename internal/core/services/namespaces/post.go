package namespaces

import (
	ports "pod-chef-back-end/internal/core/ports"
)

func NewAddService(kubernetesRepository ports.Namespace) *Service {
	return &Service{
		kubernetesRepository: kubernetesRepository,
	}
}

func (srv *Service) CreateNamespace(name string) (interface{}, error) {
	response, err := srv.kubernetesRepository.CreateNamespace(name)

	if err != nil {
		return nil, err
	}

	return response, nil
}
