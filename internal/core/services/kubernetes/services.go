package kubernetes

//GetServicesByNamespace service responsible for getting all kubernetes services inside a namespace
func (srv *Service) GetServicesByNamespace(namespace string) (interface{}, error) {
	//call driven adapter responsible for getting services from the kubernetes cluster inside a namespace
	response, err := srv.kubernetesRepository.GetServicesByNamespace(namespace)

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	return response, nil
}

//GetServiceByNameAndNamespace service responsible for getting a kubernetes service inside a namespace by it's name
func (srv *Service) GetServiceByNameAndNamespace(name string, namespace string) (interface{}, error) {
	//call driven adapter responsible for getting services from the kubernetes cluster inside a namespace
	response, err := srv.kubernetesRepository.GetServiceByNameAndNamespace(name, namespace)

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	return response, nil
}

//CreateClusterIPService service responsible for creating a kubernetes cluster ip service inside a namespace
func (srv *Service) CreateClusterIPService(namespace string, name string, containerPort int32) (interface{}, error) {
	//call driven adapter responsible for creating a kubernetes cluster ip service inside a namespace
	response, err := srv.kubernetesRepository.CreateClusterIPService(namespace, name, containerPort)

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	return response, nil
}
