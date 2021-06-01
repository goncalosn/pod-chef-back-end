package mongo

import (
	ports "pod-chef-back-end/internal/core/ports"
)

//Service Mongo repository
type Service struct {
	mongoRepository ports.MongoRepository
}

//NewMongoService where the mongo repository is injected
func NewMongoService(mongoRepository ports.MongoRepository) *Service {
	return &Service{
		mongoRepository: mongoRepository,
	}
}
