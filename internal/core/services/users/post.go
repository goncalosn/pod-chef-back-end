package users

func (srv *Service) Register(username string, email string, password string) (interface{}, error) {
	res, err := srv.mongoRepository.Register(username, email, password)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (srv *Service) Authenticate(email string, password string) (interface{}, error) {
	res, err := srv.mongoRepository.Authenticate(email, password)
	if err != nil {
		return nil, err
	}
	return res, nil
}
