package email

import "pod-chef-back-end/internal/core/ports"

//Service email repository
type Service struct {
	EmailRepository ports.EmailRepository
}

//NewEmailService where the email repository is injected
func NewEmailService(emailRepository ports.EmailRepository) *Service {
	return &Service{
		EmailRepository: emailRepository,
	}
}
