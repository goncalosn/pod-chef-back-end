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

	if response == nil { //user doesn't exist
		return nil, &pkg.Error{Err: err, Code: http.StatusNotFound, Message: "User with these credentials not found"}
	}

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	return response, nil
}

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
func (srv *Service) InsertUser(email string, hash string, name string) (*models.User, error) {
	//check if the email already exists
	response, err := srv.mongoRepository.GetUserByEmail(email)

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	var insertResponse *models.User

	if response == nil { //email is not being used
		//call driven adapter responsible for inserting a user inside the database
		insertResponse, err = srv.mongoRepository.InsertUser(email, hash, name, "member")

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
func (srv *Service) DeleteUser(id string) (bool, error) {
	//check if the email already exists
	response, err := srv.mongoRepository.GetUserByID(id)
	if response == nil { //user doesn't exist
		return false, &pkg.Error{Err: err, Code: http.StatusNotFound, Message: "User with these credentials not found"}
	}

	//call driven adapter responsible for inserting a user inside the database
	_, err = srv.mongoRepository.DeleteUserByID(id)

	if err != nil {
		//return the error sent by the repository
		return false, err
	}

	responseDeployments, err := srv.mongoRepository.GetDeploymentsFromUser(id)
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

//UpdateSelfPassword service responsible for updating a user password
func (srv *Service) UpdateSelfPassword(email string, hash string) (bool, error) {
	//check if the email already exists
	user, err := srv.mongoRepository.GetUserByEmail(email)
	if user == nil { //user doesn't exist
		return false, &pkg.Error{Err: err, Code: http.StatusNotFound, Message: "User with these credentials not found"}
	}

	//call driven adapter responsible for updating a user's password inside the database
	_, err = srv.mongoRepository.UpdateUserPassword(user.ID, hash)

	if err != nil {
		//return the error sent by the repository
		return false, err
	}

	return true, nil
}

//ResetUserPassword service responsible for updating a user password and sending an email with it
func (srv *Service) ResetUserPassword(id string, password string) (bool, error) {
	//generate hash from the password
	crypt := pkg.EncryptPassword(password)

	//check if the email already exists
	user, err := srv.mongoRepository.GetUserByID(id)
	if user == nil { //user doesn't exist
		return false, &pkg.Error{Err: err, Code: http.StatusNotFound, Message: "User with these credentials not found"}
	}

	//call driven adapter responsible for updating a user's password inside the database
	_, err = srv.mongoRepository.UpdateUserPassword(id, string(crypt))

	if err != nil {
		//return the error sent by the repository
		return false, err
	}

	data := &email.Email{
		SenderName: "Pod Chef team",
		Subject:    "Your password has been reseted!",
		Password:   password,
	}

	//call driven adapter responsible for sending an email
	_, err = srv.emailRepository.SendEmailSMTP(user.Email, data, "password-reset.txt")

	return true, nil
}

//UpdateUserRole service responsible for updating a user's role
func (srv *Service) UpdateUserRole(id string, role string) (bool, error) {
	//check if the user exists
	response, err := srv.mongoRepository.GetUserByID(id)
	if response == nil { //user doesn't exist
		return false, &pkg.Error{Err: err, Code: http.StatusNotFound, Message: "User not found"}
	}

	//call driven adapter responsible for updating a user's role inside the database
	_, err = srv.mongoRepository.UpdateUserRole(id, role)

	if err != nil {
		//return the error sent by the repository
		return false, err
	}

	return true, nil
}

//UpdateUserName service responsible for updating a user's name
func (srv *Service) UpdateUserName(id string, name string) (bool, error) {
	//check if the email already exists
	user, err := srv.mongoRepository.GetUserByID(id)
	if user == nil { //user doesn't exist
		return false, &pkg.Error{Err: err, Code: http.StatusNotFound, Message: "User not found"}
	}

	//call driven adapter responsible for updating a user's name inside the database
	_, err = srv.mongoRepository.UpdateUserName(user.Email, name)

	if err != nil {
		//return the error sent by the repository
		return false, err
	}

	return true, nil
}
