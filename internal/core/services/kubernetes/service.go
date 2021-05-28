package kubernetes

import (
	ports "pod-chef-back-end/internal/core/ports"
)

type Service struct {
	kubernetesRepository ports.KubernetesRepository
}

func NewKubernetesService(kubernetesRepository ports.KubernetesRepository) *Service {
	return &Service{
		kubernetesRepository: kubernetesRepository,
	}
}
