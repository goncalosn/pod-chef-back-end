package services

import ports "pod-chef-back-end/internal/core/ports"

func NewPostService(kubernetesRepository ports.Service) *Service {
	return &Service{
		kubernetesRepository: kubernetesRepository,
	}
}

func (srv *Service) CreateClusterIPService(namespace string, name string) (interface{}, error) {
	response, err := srv.kubernetesRepository.CreateClusterIPService(namespace, name)

	if err != nil {
		return nil, err
	}

	return response, nil
}
