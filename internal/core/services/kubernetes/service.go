package kubernetes

import (
	ports "pod-chef-back-end/internal/core/ports"
)

//Service kubernetes repository
type Service struct {
	kubernetesRepository ports.KubernetesRepository
}

//NewKubernetesService where the kubernetes repository is injected
func NewKubernetesService(kubernetesRepository ports.KubernetesRepository) *Service {
	return &Service{
		kubernetesRepository: kubernetesRepository,
	}
}
