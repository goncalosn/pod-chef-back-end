package nodes

import (
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

func (h *NodeInteractor) GetNodeStatsServiceInteractor(name string) (Node, error) {
	node, err := h.NodeService.GetNodeStatsService(name)

	result := Node{MemoryPressure: node.MemoryPressure, DiskPressure: node.DiskPressure, PIDPressure: node.PIDPressure, Ready: node.Ready}

	return result, err
}
