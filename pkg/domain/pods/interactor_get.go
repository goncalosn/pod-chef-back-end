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

func (h *PodInteractor) GetPodsByNodeAndNamespaceInteractor(node string, namespace string) ([]Pod, error) {
	result, err := h.PodService.GetPodsByNodeAndNamespaceService(node, namespace)
	var newResult []Pod

	for _, kubePod := range result {
		newPod := Pod{Name: kubePod.Name, State: kubePod.State, RestartCount: kubePod.RestartCount}
		newResult = append(newResult, newPod)
	}

	return newResult, err
}
