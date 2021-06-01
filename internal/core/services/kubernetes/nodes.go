package kubernetes

//GetNodeByName service responsible for getting a node by it's name
func (srv *Service) GetNodeByName(name string) (interface{}, error) {
	//call driven adapter responsible for getting a node from the kubernetes cluster
	node, err := srv.kubernetesRepository.GetNodeByName(name)

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	return node, nil
}

//GetNodes service responsible for getting all the nodes
func (srv *Service) GetNodes() (interface{}, error) {
	//call driven adapter responsible for getting all the nodes from the kubernetes cluster
	nodes, err := srv.kubernetesRepository.GetNodes()

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	return nodes, nil
}
