package mongo

import (
	"net/http"
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

//InsertUser service responsible for inserting a user into the database
func (srv *Service) InsertUser(email string, hash string, tokenIv string, name string, role string) (*models.User, error) {
	//check if the email already exists
	response, err := srv.mongoRepository.GetUserByEmail(email)

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	var insertResponse *models.User

	if response == nil { //email is not being used
		//call driven adapter responsible for inserting a user inside the database
		insertResponse, err = srv.mongoRepository.InsertUser(email, hash, tokenIv, name, role)

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
func (srv *Service) DeleteUser(email string) (interface{}, error) {
	//check if the email already exists
	response, err := srv.mongoRepository.GetUserByEmail(email)
	if response == nil { //user doesn't exist
		return nil, &pkg.Error{Err: err, Code: http.StatusNotFound, Message: "User with these credentials not found"}
	}

	//call driven adapter responsible for inserting a user inside the database
	responseDelete, err := srv.mongoRepository.DeleteUserByEmail(email)

	if err != nil {
		//return the error sent by the repository
		return nil, err
	}

	return responseDelete, nil
}
