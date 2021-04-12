package pods

import (
	. "pod-chef-back-end/pkg/kubernetes/pods"
)

type PodInteractor struct {
	PodService *PodService
}

func NewPodInteractor(PodService *PodService) *PodInteractor {
	return &PodInteractor{
		PodService: PodService,
	}
}

func (h *PodInteractor) GetPodsByNodeAndNamespaceInteractor(node string, namespace string) (interface{}, error) {
	pods, err := h.PodService.GetPodsByNodeAndNamespaceService(node, namespace)

	return pods, err
}
