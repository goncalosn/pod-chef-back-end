package nodes

import (
	. "pod-chef-back-end/pkg/domain/pods"
	. "pod-chef-back-end/pkg/kubernetes/nodes"
)

type NodeInteractor struct {
	NodeService *NodeService
}

func NewNodeInteractor(NodeService *NodeService) *NodeInteractor {
	return &NodeInteractor{
		NodeService: NodeService,
	}
}

func (h *NodeInteractor) GetNodeStatsServiceInteractor(node string, namespace string) ([]Pod, error) {
	result, err := h.NodeService.GetNodeStatsService(node, namespace)
	var newResult []Pod

	for _, kubePod := range result {
		newPod := Pod{Name: kubePod.Name, State: kubePod.State, RestartCount: kubePod.RestartCount}
		newResult = append(newResult, newPod)
	}

	return newResult, err
}
