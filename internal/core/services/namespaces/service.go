package namespaces

import "pod-chef-back-end/internal/core/ports"

type Service struct {
	kubernetesRepository ports.Namespace
}

func NewService(kubernetesRepository ports.Namespace) *Service {
	return &Service{
		kubernetesRepository: kubernetesRepository,
	}
}
