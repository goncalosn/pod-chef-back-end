package kubernetes

//CreateNamespace service responsible for creating a namespace inside the kubernetes cluster
func (srv *Service) CreateNamespace(name string) (interface{}, error) {
	//call driven adapter responsible for creating a namespace inside the kubernetes cluster
	response, err := srv.kubernetesRepository.CreateNamespace(name)

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	return response, nil
}

//DeleteNamespace service responsible for deleting a namespace form the kubernetes cluster
func (srv *Service) DeleteNamespace(name string) (interface{}, error) {
	//call driven adapter responsible for deleting a namespace from the kubernetes cluster
	response, err := srv.kubernetesRepository.DeleteNamespace(name)

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	return response, nil
}
