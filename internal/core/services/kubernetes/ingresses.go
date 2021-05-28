package kubernetes

//GetIngressByName service responsible for getting an ingress inside a namespace by it's name
func (srv *Service) GetIngressByName(name string, namespace string) (interface{}, error) {
	//call driven adapter responsible for getting the ingress inside the kubernetes cluster
	response, err := srv.kubernetesRepository.GetIngressByNameAndNamespace(name, namespace)

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	return response, nil
}

//CreateIngress service responsible for creating an ingress inside a namespace
func (srv *Service) CreateIngress(namespace string, name string, host string) (interface{}, error) {
	//call driven adapter responsible for creating an ingress inside the kubernetes cluster
	response, err := srv.kubernetesRepository.CreateIngress(namespace, name, host)

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	return response, nil
}
