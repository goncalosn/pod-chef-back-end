package services

import ports "pod-chef-back-end/internal/core/ports"

type Service struct {
	kubernetesRepository ports.Volume
}

func NewService(kubernetesRepository ports.Volume) *Service {
	return &Service{
		kubernetesRepository: kubernetesRepository,
	}
}
