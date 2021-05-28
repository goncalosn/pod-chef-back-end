package kubernetes

func (srv *Service) GetNamespaces() (interface{}, error) {
	response, err := srv.kubernetesRepository.GetNamespaces()

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (srv *Service) CreateNamespace(name string) (interface{}, error) {
	response, err := srv.kubernetesRepository.CreateNamespace(name)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (srv *Service) DeleteNamespace(name string) (interface{}, error) {
	response, err := srv.kubernetesRepository.DeleteNamespace(name)

	if err != nil {
		return nil, err
	}

	return response, nil
}
