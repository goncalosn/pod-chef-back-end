package deployments

import "pod-chef-back-end/internal/core/ports"

type Service struct {
	kubernetesRepository ports.Deployment
}

func NewService(kubernetesRepository ports.Deployment) *Service {
	return &Service{
		kubernetesRepository: kubernetesRepository,
	}
}
