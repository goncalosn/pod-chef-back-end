package services

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
