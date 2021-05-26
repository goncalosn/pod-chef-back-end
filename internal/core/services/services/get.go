package services

import ports "pod-chef-back-end/internal/core/ports"

func NewGetService(kubernetesRepository ports.Service) *Service {
	return &Service{
		kubernetesRepository: kubernetesRepository,
	}
}

func (srv *Service) GetServicesByNamespace(namespace string) (interface{}, error) {
	servs, err := srv.kubernetesRepository.GetServicesByNamespace(namespace)

	if err != nil {
		return nil, err
	}

	return servs, nil
}

func (srv *Service) GetServiceByNameAndNamespace(name string, namespace string) (interface{}, error) {
	servs, err := srv.kubernetesRepository.GetServiceByNameAndNamespace(name, namespace)

	if err != nil {
		return nil, err
	}

	return servs, nil
}
