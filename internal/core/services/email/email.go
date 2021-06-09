package email

import (
	email "pod-chef-back-end/internal/core/domain/email"
)

//SendEmail service responsible for sending a new email
func (srv *Service) SendEmail(to string, subject string, template string) (interface{}, error) {

	data := &email.Email{
		ReceiverName: "David Gilmour",
		SenderName:   "Pod Chef team",
		Subject:      subject,
	}

	//call driven adapter responsible for sending an email
	response, err := srv.EmailRepository.SendEmailSMTP(to, data, template)

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	return response, nil
}
