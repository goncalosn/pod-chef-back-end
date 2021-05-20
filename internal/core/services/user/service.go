package user

import ports "pod-chef-back-end/internal/core/ports"

type Service struct {
	mongoRepository ports.User
}

func NewService(mongoRepository ports.User) *Service {
	return &Service{
		mongoRepository: mongoRepository,
	}
}
