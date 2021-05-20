package user

import (
	"pod-chef-back-end/internal/core/ports"
	"pod-chef-back-end/pkg/auth"
)

func NewPostService(mongoRepository ports.User) *Service {
	return &Service{
		mongoRepository: mongoRepository,
	}
}

func (srv *Service) Login(email string) (interface{}, error) {
	res, err := srv.mongoRepository.Login(email)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (srv *Service) SignIn(user auth.User) (interface{}, error) {
	res, err := srv.mongoRepository.SignIn(user)
	if err != nil {
		return nil, err
	}
	return res, nil
}
