package kubernetes

//GetPodsByNodeAndNamespace service responsible for getting all pods inside a namespace and node
func (srv *Service) GetPodsByNodeAndNamespace(node string, namespace string) (interface{}, error) {
	//call driven adapter responsible for for getting all pods inside a namespace and node
	pods, err := srv.kubernetesRepository.GetPodsByNodeAndNamespace(node, namespace)

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	return pods, nil
}
