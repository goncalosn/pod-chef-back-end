package pods

import (
	"errors"
	ports "pod-chef-back-end/internal/core/ports"
)

type service struct {
	kubernetesRepository ports.Pod
}

func NewGetService(kubernetesRepository ports.Pod) *service {
	return &service{
		kubernetesRepository: kubernetesRepository,
	}
}

func (srv *service) GetPodsByNodeAndNamespace(node string, namespace string) (interface{}, error) {
	pods, err := srv.kubernetesRepository.GetPodsByNodeAndNamespace(node, namespace)

	if err != nil {
		return nil, errors.New("")
	}

	return pods, nil
}
