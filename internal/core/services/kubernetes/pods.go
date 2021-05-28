package kubernetes

func (srv *Service) GetPodsByNodeAndNamespace(node string, namespace string) (interface{}, error) {
	pods, err := srv.kubernetesRepository.GetPodsByNodeAndNamespace(node, namespace)

	if err != nil {
		return nil, err
	}

	return pods, nil
}
