package email

import (
	"net/http"
	email "pod-chef-back-end/internal/core/domain/email"
	"pod-chef-back-end/pkg"
)

//SendInvitationEmail service responsible for sending a new invitation email
func (srv *Service) SendInvitationEmail(to string, subject string, template string) (interface{}, error) {

	data := &email.Email{
		SenderName: "Pod Chef team",
		Subject:    subject,
	}

	//verify if user exists
	user, err := srv.mongoRepository.GetUserFromWhitelistByEmail(to)
	if err != nil {
		return nil, err
	}

	if user != nil {
		//return a custom error
		return nil, &pkg.Error{Err: err, Code: http.StatusNotFound, Message: "User already invited"}
	}

	//add user to whitelist
	_, err = srv.mongoRepository.InsertUserIntoWhitelist(to)
	if err != nil {
		return nil, err
	}

	//call driven adapter responsible for sending an email
	response, err := srv.emailRepository.SendEmailSMTP(to, data, template)

	if err != nil { //transaction
		//delete user from whitelist
		_, err = srv.mongoRepository.DeleteUserFromWhitelistByEmail(to)

		//return the error sent by the repository
		return nil, err
	}

	return response, nil
}

//SendAnnulmentEmail service responsible for sending a new annulment email
func (srv *Service) SendAnnulmentEmail(to string, subject string, template string) (interface{}, error) {

	data := &email.Email{
		SenderName: "Pod Chef team",
		Subject:    subject,
	}

	//verify if user exists
	user, err := srv.mongoRepository.GetUserFromWhitelistByEmail(to)
	if err != nil {
		return nil, err
	}

	if user == nil {
		//return a custom error
		return nil, &pkg.Error{Err: err, Code: http.StatusNotFound, Message: "User not found"}
	}

	//add delete user from whitelist
	_, err = srv.mongoRepository.DeleteUserFromWhitelistByEmail(to)
	if err != nil {
		return nil, err
	}

	//call driven adapter responsible for sending an email
	response, err := srv.emailRepository.SendEmailSMTP(to, data, template)

	if err != nil { //transaction
		//delete user from whitelist
		_, err = srv.mongoRepository.DeleteUserFromWhitelistByEmail(to)

		//return the error sent by the repository
		return nil, err
	}

	return response, nil
}
