package services

func (srv *Service) GetServicesByNamespace(namespace string) (interface{}, error) {
	//TODO: check for namespace
	servs, err := srv.kubernetesRepository.GetServicesByNamespace(namespace)

	if err != nil {
		return nil, err
	}

	return servs, nil
}
