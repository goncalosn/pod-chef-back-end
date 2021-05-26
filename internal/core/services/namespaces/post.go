package namespaces

import (
	ports "pod-chef-back-end/internal/core/ports"
)

func NewAddService(kubernetesRepository ports.Namespace) *Service {
	return &Service{
		kubernetesRepository: kubernetesRepository,
	}
}

func (srv *Service) AddNamespace(name string) (interface{}, error) {
	response, err := srv.kubernetesRepository.AddNamespace(name)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (srv *Service) CheckRepeatedNamespace(name string) (interface{}, error) {
	response, err := srv.kubernetesRepository.CheckRepeatedNamespace(name)

	if err != nil {
		return nil, err
	}

	return response, nil
}
