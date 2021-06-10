package mongo

import (
	"net/http"
	email "pod-chef-back-end/internal/core/domain/email"
	models "pod-chef-back-end/internal/core/domain/mongo"
	"pod-chef-back-end/pkg"
)

//GetUserByEmail service responsible for getting a user from the database
func (srv *Service) GetUserByEmail(email string) (*models.User, error) {
	//call driven adapter responsible for getting a deployment from mongo database
	response, err := srv.mongoRepository.GetUserByEmail(email)

	if response == nil { //user doesn't exist
		return nil, &pkg.Error{Err: err, Code: http.StatusNotFound, Message: "User with these credentials not found"}
	}

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	return response, nil
}

//GetAllUsers service responsible for getting all users from the database
func (srv *Service) GetAllUsers() (*[]models.User, error) {
	//call driven adapter responsible for getting a deployment from mongo database
	response, err := srv.mongoRepository.GetAllUsers()

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	return response, nil
}

//InsertUser service responsible for inserting a user into the database
func (srv *Service) InsertUser(email string, hash string, name string, role string) (*models.User, error) {
	//check if the email already exists
	response, err := srv.mongoRepository.GetUserByEmail(email)

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	var insertResponse *models.User

	if response == nil { //email is not being used
		//call driven adapter responsible for inserting a user inside the database
		insertResponse, err = srv.mongoRepository.InsertUser(email, hash, name, role)

		if err != nil {
			//return the error sent by the repository
			return nil, err
		}

		//delete user invitation
		_, err := srv.mongoRepository.DeleteUserFromWhitelistByEmail(email)

		if err != nil {
			//return the error sent by the repository
			return nil, err
		}
	} else { //email being used already
		//return a custom error
		return nil, &pkg.Error{Err: err, Code: http.StatusBadRequest, Message: "Email already in use"}
	}

	return insertResponse, nil
}

//DeleteUser service responsible for deleting a user from the database
func (srv *Service) DeleteUser(email string) (bool, error) {
	//check if the email already exists
	response, err := srv.mongoRepository.GetUserByEmail(email)
	if response == nil { //user doesn't exist
		return false, &pkg.Error{Err: err, Code: http.StatusNotFound, Message: "User with these credentials not found"}
	}

	//call driven adapter responsible for inserting a user inside the database
	_, err = srv.mongoRepository.DeleteUserByEmail(email)

	if err != nil {
		//return the error sent by the repository
		return false, err
	}

	responseDeployments, err := srv.mongoRepository.GetDeploymentsFromUser(email)
	if err != nil {
		//return the error sent by the repository
		return false, err
	}

	for _, element := range responseDeployments {
		_, err := srv.mongoRepository.DeleteDeploymentByUUID(element.UUID)
		if err != nil {
			//return the error sent by the repository
			return false, err
		}
	}

	return true, nil
}

//GetAllUsersFromWhitelist service responsible for getting all users from the whitelist
func (srv *Service) GetAllUsersFromWhitelist() (*[]models.User, error) {
	//call driven adapter responsible for getting a deployment from mongo database
	response, err := srv.mongoRepository.GetAllUsersFromWhitelist()

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	return &response, nil
}

//InviteUserToWhitelist service responsible for deleting a user from the database
func (srv *Service) InviteUserToWhitelist(to string) (bool, error) {
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
	_, err = srv.mongoRepository.InsertUserIntoWhitelist(to)
	if err != nil {
		return false, err
	}

	//call driven adapter responsible for sending an email
	response, err := srv.emailRepository.SendEmailSMTP(to, data, "invitation.txt")

	if err != nil { //transaction
		//delete user from whitelist
		_, err = srv.mongoRepository.DeleteUserFromWhitelistByEmail(to)

		//return the error sent by the repository
		return false, err
	}

	return response, nil
}

//RemoveUserFromWhitelist service responsible for deleting a user from the database
func (srv *Service) RemoveUserFromWhitelist(to string) (bool, error) {
	data := &email.Email{
		SenderName: "Pod Chef team",
		Subject:    "You have been removed from the Pod Chef whitelist!",
	}

	//verify if user exists
	user, err := srv.mongoRepository.GetUserFromWhitelistByEmail(to)
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
	_, err = srv.mongoRepository.DeleteUserFromWhitelistByEmail(to)
	if err != nil {
		return false, err
	}

	//call driven adapter responsible for sending an email
	response, err := srv.emailRepository.SendEmailSMTP(to, data, "annulment.txt")

	if err != nil { //delete user from whitelist
		//delete user from whitelist
		_, err = srv.mongoRepository.DeleteUserFromWhitelistByEmail(to)

		//return the error sent by the repository
		return false, err
	}

	return response, nil
}
