package pkg

import (
	nodes "pod-chef-back-end/pkg/domain/nodes"
	pods "pod-chef-back-end/pkg/domain/pods"
)

type InteractorsContainer struct {
	NodeInteractor *nodes.NodeInteractor
	PodInteractor  *pods.PodInteractor
}

func BuildInteractors(s *ServicesContainer) *InteractorsContainer {
	return &InteractorsContainer{
		NodeInteractor: nodes.NewNodeInteractor(s.NodeService),
		PodInteractor:  pods.NewPodInteractor(s.PodService),
	}
}
