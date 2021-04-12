package nodes

import (
	"errors"
	ports "pod-chef-back-end/internal/core/ports"
)

type service struct {
	kubernetesRepository ports.Node
}

func NewGetService(kubernetesRepository ports.Node) *service {
	return &service{
		kubernetesRepository: kubernetesRepository,
	}
}

func (srv *service) GetNode(name string) (interface{}, error) {
	node, err := srv.kubernetesRepository.GetNode(name)

	if err != nil {
		return nil, errors.New("")
	}

	return node, nil
}

func (srv *service) GetNodes() (interface{}, error) {
	nodes, err := srv.kubernetesRepository.GetNodes()

	if err != nil {
		return nil, errors.New("")
	}

	return nodes, nil
}
