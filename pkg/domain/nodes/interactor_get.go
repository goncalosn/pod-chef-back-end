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

func (h *NodeInteractor) GetNodeInteractor(name string) (interface{}, error) {
	node, err := h.NodeService.GetNodeService(name)

	return node, err
}

func (h *NodeInteractor) GetNodesInteractor() (interface{}, error) {
	nodes, err := h.NodeService.GetNodesService()

	return nodes, err
}
