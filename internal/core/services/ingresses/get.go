package ingresses

import (
	ports "pod-chef-back-end/internal/core/ports"
)

func NewGetService(kubernetesRepository ports.Ingress) *Service {
	return &Service{
		kubernetesRepository: kubernetesRepository,
	}
}

func (srv *Service) GetIngress(name string) (interface{}, error) {
	response, err := srv.kubernetesRepository.GetIngress(name)

	if err != nil {
		return nil, err
	}

	return response, nil
}
