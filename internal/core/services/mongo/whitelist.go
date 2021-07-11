package mongo

import (
	"net/http"
	email "pod-chef-back-end/internal/core/domain/email"
	models "pod-chef-back-end/internal/core/domain/mongo"
	"pod-chef-back-end/pkg"
)

//GetAllUsersFromWhitelist service responsible for getting all users from the whitelist
func (srv *Service) GetAllUsersFromWhitelist() (*[]models.WhitelistUser, error) {
	//call driven adapter responsible for getting a deployment from mongo database
	response, err := srv.mongoRepository.GetAllUsersFromWhitelist()

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	return &response, nil
}

//InsertUserIntoWhitelist service responsible for deleting a user from the database
func (srv *Service) InsertUserIntoWhitelist(to string) (*string, error) {
	//verify if user exists
	user, err := srv.mongoRepository.GetUserByEmail(to)
	if err != nil {
		if mongoError := err.(*pkg.Error); mongoError.Code != http.StatusNotFound {
			return nil, err
		}
	}

	if user != nil {
		//return a custom error
		return nil, &pkg.Error{Err: err, Code: http.StatusNotFound, Message: "A user with this email already exists"}
	}

	//verify if user is already invited
	invited, err := srv.mongoRepository.GetUserFromWhitelistByEmail(to)
	if err != nil {
		if mongoError := err.(*pkg.Error); mongoError.Code != http.StatusNotFound {
			return nil, err
		}
	}

	if invited != nil {
		//return a custom error
		return nil, &pkg.Error{Err: err, Code: http.StatusNotFound, Message: "User already invited"}
	}

	//add user to whitelist
	_, err = srv.mongoRepository.InsertUserIntoWhitelist(to)
	if err != nil {
		return nil, err
	}

	data := &email.Email{
		SenderName: "Pod Chef team",
		Subject:    "You have been added to the Pod Chef whitelist!",
	}

	emailBody :=
		"Hi,\n" +
			"You are now in the whitelist of the Pod Chef project and can register https://wwww.podchef.cf\n" +
			"\nWith The best regards, " + data.SenderName + "."

	//call driven adapter responsible for sending an email
	err = srv.emailRepository.SendEmailSMTP(to, data.Subject, emailBody)

	message := "User invited sucessfully"

	return &message, nil
}

//RemoveUserFromWhitelist service responsible for deleting a user from the database
func (srv *Service) RemoveUserFromWhitelist(id string) (*string, error) {

	//verify if user exists
	user, err := srv.mongoRepository.GetUserFromWhitelistByID(id)
	if err != nil {
		return nil, err
	}

	//delete user from whitelist
	_, err = srv.mongoRepository.DeleteUserFromWhitelistByID(id)
	if err != nil {
		return nil, err
	}

	data := &email.Email{
		SenderName: "Pod Chef team",
		Subject:    "You have been removed from the Pod Chef whitelist!",
	}

	emailBody :=
		"Hi,\n" +
			"You have been removed from the Pod Chef whitelist.\n" +
			"\nWith The best regards, " + data.SenderName + "."

	//call driven adapter responsible for sending an email
	err = srv.emailRepository.SendEmailSMTP(user.Email, data.Subject, emailBody)

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	message := "User removed from whitelist"

	return &message, nil
}
