package nodes

import ports "pod-chef-back-end/internal/core/ports"

type Service struct {
	kubernetesRepository ports.Node
}

func NewService(kubernetesRepository ports.Node) *Service {
	return &Service{
		kubernetesRepository: kubernetesRepository,
	}
}
