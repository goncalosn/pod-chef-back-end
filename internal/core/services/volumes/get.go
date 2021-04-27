package services

func (srv *Service) GetVolumes() (interface{}, error) {
	servs, err := srv.kubernetesRepository.GetVolumes()
	if err != nil {
		return nil, err
	}
	return servs, nil
}
