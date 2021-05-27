package ingresses

import (
	ports "pod-chef-back-end/internal/core/ports"
)

func NewPostService(kubernetesRepository ports.Ingress) *Service {
	return &Service{
		kubernetesRepository: kubernetesRepository,
	}
}

func (srv *Service) CreateIngress(namespace string, name string, host string) (interface{}, error) {
	response, err := srv.kubernetesRepository.CreateIngress(namespace, name, host)

	if err != nil {
		return nil, err
	}

	return response, nil
}
