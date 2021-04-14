package pods

import ports "pod-chef-back-end/internal/core/ports"

type Service struct {
	kubernetesRepository ports.Pod
}

func NewService(kubernetesRepository ports.Pod) *Service {
	return &Service{
		kubernetesRepository: kubernetesRepository,
	}
}
