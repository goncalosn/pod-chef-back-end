package kubernetes

func (srv *Service) GetServicesByNamespace(namespace string) (interface{}, error) {
	servs, err := srv.kubernetesRepository.GetServicesByNamespace(namespace)

	if err != nil {
		return nil, err
	}

	return servs, nil
}

func (srv *Service) GetServiceByNameAndNamespace(name string, namespace string) (interface{}, error) {
	servs, err := srv.kubernetesRepository.GetServiceByNameAndNamespace(name, namespace)

	if err != nil {
		return nil, err
	}

	return servs, nil
}

func (srv *Service) CreateClusterIPService(namespace string, name string) (interface{}, error) {
	response, err := srv.kubernetesRepository.CreateClusterIPService(namespace, name)

	if err != nil {
		return nil, err
	}

	return response, nil
}
