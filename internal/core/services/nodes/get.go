package nodes

import ports "pod-chef-back-end/internal/core/ports"

func NewGetService(kubernetesRepository ports.Node) *Service {
	return &Service{
		kubernetesRepository: kubernetesRepository,
	}
}

func (srv *Service) GetNode(name string) (interface{}, error) {
	node, err := srv.kubernetesRepository.GetNode(name)

	if err != nil {
		return nil, err
	}

	return node, nil
}

func (srv *Service) GetNodes() (interface{}, error) {
	nodes, err := srv.kubernetesRepository.GetNodes()

	if err != nil {
		return nil, err
	}

	return nodes, nil
}
