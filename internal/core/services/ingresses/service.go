package ingresses

import "pod-chef-back-end/internal/core/ports"

type Service struct {
	kubernetesRepository ports.Ingress
}

func NewService(kubernetesRepository ports.Ingress) *Service {
	return &Service{
		kubernetesRepository: kubernetesRepository,
	}
}
