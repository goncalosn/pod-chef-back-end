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
func (srv *Service) InsertUserIntoWhitelist(to string) (bool, error) {
	data := &email.Email{
		SenderName: "Pod Chef team",
		Subject:    "You have been added to the Pod Chef whitelist!",
	}

	//verify if user exists
	user, err := srv.mongoRepository.GetUserFromWhitelistByEmail(to)
	if err != nil {
		mongoError := err.(*pkg.Error)
		if mongoError.Code != http.StatusNotFound {
			return false, err
		}
	}

	if user != nil {
		//return a custom error
		return false, &pkg.Error{Err: err, Code: http.StatusNotFound, Message: "User already invited"}
	}

	//add user to whitelist
	id, err := srv.mongoRepository.InsertUserIntoWhitelist(to)
	if err != nil {
		return false, err
	}

	//call driven adapter responsible for sending an email
	response, err := srv.emailRepository.SendEmailSMTP(to, data, "invitation.txt")

	if err != nil { //transaction
		//delete user from whitelist
		_, err = srv.mongoRepository.DeleteUserFromWhitelistByID(*id)

		//return the error sent by the repository
		return false, err
	}

	return response, nil
}

//RemoveUserFromWhitelist service responsible for deleting a user from the database
func (srv *Service) RemoveUserFromWhitelist(id string) (bool, error) {
	data := &email.Email{
		SenderName: "Pod Chef team",
		Subject:    "You have been removed from the Pod Chef whitelist!",
	}

	//verify if user exists
	user, err := srv.mongoRepository.GetUserFromWhitelistByID(id)
	if err != nil {
		mongoError := err.(*pkg.Error)
		if mongoError.Code != http.StatusNotFound {
			return false, err
		}
	}

	if user == nil {
		//return a custom error
		return false, &pkg.Error{Err: err, Code: http.StatusNotFound, Message: "User not found"}
	}

	//delete user from whitelist
	_, err = srv.mongoRepository.DeleteUserFromWhitelistByID(id)
	if err != nil {
		return false, err
	}

	//call driven adapter responsible for sending an email
	response, err := srv.emailRepository.SendEmailSMTP(user.Email, data, "annulment.txt")

	if err != nil {
		//return the error sent by the repository
		return false, err
	}

	return response, nil
}
