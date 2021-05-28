package kubernetes

func (srv *Service) GetIngress(name string) (interface{}, error) {
	response, err := srv.kubernetesRepository.GetIngress(name)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (srv *Service) CreateIngress(namespace string, name string, host string) (interface{}, error) {
	response, err := srv.kubernetesRepository.CreateIngress(namespace, name, host)

	if err != nil {
		return nil, err
	}

	return response, nil
}
