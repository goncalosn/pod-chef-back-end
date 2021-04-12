package nodes

import (
	"errors"
	ports "pod-chef-back-end/internal/core/ports"
)

type Service struct {
	kubernetesRepository ports.Node
}

func NewGetService(kubernetesRepository ports.Node) *Service {
	return &Service{
		kubernetesRepository: kubernetesRepository,
	}
}

func (srv *Service) GetNode(name string) (interface{}, error) {
	node, err := srv.kubernetesRepository.GetNode(name)

	if err != nil {
		return nil, errors.New("")
	}

	return node, nil
}

func (srv *Service) GetNodes() (interface{}, error) {
	nodes, err := srv.kubernetesRepository.GetNodes()

	if err != nil {
		return nil, errors.New("")
	}

	return nodes, nil
}
