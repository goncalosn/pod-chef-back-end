package mongo

import (
	"net/http"
	email "pod-chef-back-end/internal/core/domain/email"
	models "pod-chef-back-end/internal/core/domain/mongo"
	"pod-chef-back-end/pkg"
)

//GetUserByID service responsible for getting a user from the database
func (srv *Service) GetUserByID(id string) (*models.User, error) {
	//call driven adapter responsible for getting a deployment from mongo database
	response, err := srv.mongoRepository.GetUserByID(id)

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	if response == nil { //user doesn't exist
		return nil, &pkg.Error{Err: err, Code: http.StatusNotFound, Message: "User with these credentials not found"}
	}

	return response, nil
}

//GetUserByEmail service responsible for getting a user from the database
func (srv *Service) GetUserByEmail(email string) (*models.User, error) {
	//call driven adapter responsible for getting a deployment from mongo database
	response, err := srv.mongoRepository.GetUserByEmail(email)

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	if response == nil { //user doesn't exist
		return nil, &pkg.Error{Err: err, Code: http.StatusNotFound, Message: "User with these credentials not found"}
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
func (srv *Service) InsertUser(email string, hash string, name string) (*models.User, error) {
	//check if the email already exists
	_, err := srv.mongoRepository.GetUserByEmail(email)

	if err != nil {
		if mongoError := err.(*pkg.Error); mongoError.Code != http.StatusNotFound {
			return nil, err
		}
	}

	//check if the email already exists
	_, err = srv.mongoRepository.GetUserFromWhitelistByEmail(email)

	if err != nil {
		return nil, err
	}

	var user *models.User

	//call driven adapter responsible for inserting a user inside the database
	user, err = srv.mongoRepository.InsertUser(email, hash, name, "member")

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	//remove user from whitelist
	_, err = srv.mongoRepository.DeleteUserFromWhitelistByEmail(email)

	if err != nil {
		return nil, err
	}

	return user, nil
}

//DeleteUser service responsible for deleting a user from the database
func (srv *Service) DeleteUser(id string) (*string, error) {
	//check if the email already exists
	response, err := srv.mongoRepository.GetUserByID(id)
	if response == nil { //user doesn't exist
		return nil, &pkg.Error{Err: err, Code: http.StatusNotFound, Message: "User not found"}
	}

	//call driven adapter responsible for inserting a user inside the database
	_, err = srv.mongoRepository.DeleteUserByID(id)
	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	responseDeployments, err := srv.mongoRepository.GetDeploymentsFromUser(id)
	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	for _, element := range responseDeployments {
		_, err := srv.mongoRepository.DeleteDeploymentByUUID(element.UUID)
		if err != nil {
			//return the error sent by the repository
			return nil, err
		}
	}

	message := "User deleted sucessfully"

	return &message, nil
}

//UpdateUserPassword service responsible for updating a user password and sending an email with it
func (srv *Service) UpdateUserPassword(id string, password string) (*string, error) {
	//generate hash from the password
	crypt := pkg.EncryptPassword(password)

	//check if the email already exists
	user, err := srv.mongoRepository.GetUserByID(id)
	if user == nil { //user doesn't exist
		return nil, &pkg.Error{Err: err, Code: http.StatusNotFound, Message: "User with these credentials not found"}
	}

	//call driven adapter responsible for updating a user's password inside the database
	_, err = srv.mongoRepository.UpdateUserPassword(id, string(crypt))

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	message := "Password updated sucessfully"

	return &message, nil
}

//ResetUserPassword service responsible for updating a user password and sending an email with it
func (srv *Service) ResetUserPassword(id string, password string) (*string, error) {
	//generate hash from the password
	crypt := pkg.EncryptPassword(password)

	//check if the email already exists
	user, err := srv.mongoRepository.GetUserByID(id)
	if user == nil { //user doesn't exist
		return nil, &pkg.Error{Err: err, Code: http.StatusNotFound, Message: "User with these credentials not found"}
	}

	//call driven adapter responsible for updating a user's password inside the database
	_, err = srv.mongoRepository.UpdateUserPassword(id, string(crypt))

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	data := &email.Email{
		SenderName: "Pod Chef team",
		Subject:    "Your password has been reseted!",
		Password:   password,
	}

	emailBody :=
		"Hi,\n" +
			"Your account password has been reseted.\n" +
			"Please use the following password to login next time: " + data.Password + "\n" +
			"\nWith The best regards, " + data.SenderName + ".\n" +
			"\nPS: This password should only be temporary!!!"

	//call driven adapter responsible for sending an email
	err = srv.emailRepository.SendEmailSMTP(user.Email, data.Subject, emailBody)

	message := "User's password reseted sucessfully"

	return &message, nil
}

//UpdateUserRole service responsible for updating a user's role
func (srv *Service) UpdateUserRole(id string, role string) (*string, error) {
	//check if the user exists
	response, err := srv.mongoRepository.GetUserByID(id)
	if response == nil { //user doesn't exist
		return nil, &pkg.Error{Err: err, Code: http.StatusNotFound, Message: "User not found"}
	}

	//call driven adapter responsible for updating a user's role inside the database
	_, err = srv.mongoRepository.UpdateUserRole(id, role)

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	message := "User's role updated sucessfully"

	return &message, nil
}

//UpdateUserName service responsible for updating a user's name
func (srv *Service) UpdateUserName(id string, name string) (*string, error) {
	//check if the email already exists
	_, err := srv.mongoRepository.GetUserByID(id)

	if err != nil {
		return nil, err
	}

	//call driven adapter responsible for updating a user's name inside the database
	_, err = srv.mongoRepository.UpdateUserName(id, name)

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	message := "Name updated sucessfully"

	return &message, nil
}
