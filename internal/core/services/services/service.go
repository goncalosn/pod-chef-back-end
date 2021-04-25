package services

import ports "pod-chef-back-end/internal/core/ports"

type Service struct {
	kubernetesRepository ports.Service
}

func NewService(kubernetesRepository ports.Service) *Service {
	return &Service{
		kubernetesRepository: kubernetesRepository,
	}
}
