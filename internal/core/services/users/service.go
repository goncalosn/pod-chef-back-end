package users

import "pod-chef-back-end/internal/core/ports"

type Service struct {
	mongoRepository ports.UserAuth
}

func NewService(mongoRepository ports.UserAuth) *Service {
	return &Service{
		mongoRepository: mongoRepository,
	}
}
