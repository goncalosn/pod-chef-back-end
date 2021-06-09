package email

import "pod-chef-back-end/internal/core/ports"

//Service repositories ncessary for the email methods
type Service struct {
	emailRepository ports.EmailRepository
	mongoRepository ports.MongoRepository
}

//NewEmailService where the email repository is injected
func NewEmailService(emailRepository ports.EmailRepository, mongoRepository ports.MongoRepository) *Service {
	return &Service{
		emailRepository: emailRepository,
		mongoRepository: mongoRepository,
	}
}
