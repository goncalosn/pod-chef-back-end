package pods

import ports "pod-chef-back-end/internal/core/ports"

func NewGetService(kubernetesRepository ports.Pod) *Service {
	return &Service{
		kubernetesRepository: kubernetesRepository,
	}
}

func (srv *Service) GetPodsByNodeAndNamespace(node string, namespace string) (interface{}, error) {
	pods, err := srv.kubernetesRepository.GetPodsByNodeAndNamespace(node, namespace)

	if err != nil {
		return nil, err
	}

	return pods, nil
}
