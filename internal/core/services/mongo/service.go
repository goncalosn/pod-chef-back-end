package mongo

import (
	ports "pod-chef-back-end/internal/core/ports"
)

//Service Mongo repository
type Service struct {
	mongoRepository ports.MongoRepository
	emailRepository ports.EmailRepository
}

//NewMongoService where the mongo repository is injected
func NewMongoService(mongoRepository ports.MongoRepository, emailRepository ports.EmailRepository) *Service {
	return &Service{
		mongoRepository: mongoRepository,
		emailRepository: emailRepository,
	}
}
