package service

import (
	database "lifesaver/internal/lifesaver/repo/postgres"
	"lifesaver/pkg/models"
	"github.com/google/uuid"
)

var userRepository database.UserRepository
var user database.User

const (
	DEFAULT_ERROR_IS_EXISTS		= "Unable to check for existing user"
	DEFAULT_ERROR_CREATE_USER	= "Unable to create new user."
	DEFAULT_ERROR_GET_USER		= "Unable to fetch user details for existing user."
	DEFAULT_ERROR_UPDATE_USER	= "Unable to update details for existing user."
	DEFAULT_ERROR_REMOVE_USER	= "Unable to delete existing user."
)

func IsExists(userId string) (bool, error) {
	userRepository = &user

	//Check for existing user
	isUserExists, err := userRepository.IsExists(userId)

	if err != nil {
		return false, err
	}
	if isUserExists {
		return true, nil
	}

	return false, nil;
}

func CreateUser(userToCreate *models.User) (string, *models.ErrorResponse) {
	//Generate UUID
	newUserId := uuid.New().String()

	//Validate New User Details before persiting
	validationMessage := userToCreate.ValidateInsert()
	if len(validationMessage) > 1 {
		return "", getErrorResponse(400, validationMessage)
	}

	//If there are pk constraints on columns, we should check if any records exist with same value
	//Throw apt error if such a record is found

	//Save New User
	user := database.User{Id: newUserId, Email: userToCreate.Email, Name: userToCreate.Name}
	newUserId, err := user.Save()
	if err != nil {
		return "", getErrorResponse(500, DEFAULT_ERROR_CREATE_USER)
	}

	return newUserId, nil
} 

func GetUser(userId string) (*models.User, *models.ErrorResponse) {
	userRepository = &user

	//Get Existing User
	user, err := userRepository.GetById(userId)
	if err != nil {
		return nil, getErrorResponse(500, DEFAULT_ERROR_GET_USER)
	}
	if user == nil {
		return nil, getErrorResponse(500, "No such user found.")
	}

	return user, nil
}

func UpdateUser(userToUpdate *models.User) *models.ErrorResponse {
	//Validate User Details
	validationMessage := userToUpdate.ValidateInsert()
	if len(validationMessage) > 1 {
		return getErrorResponse(400, validationMessage)
	}
	
	//Check if user already exists
	isUserExists, err := IsExists(userToUpdate.Id)
	if err != nil {
		return getErrorResponse(500, "Failed to get user datils. Unable to update user.")
	}
	if !isUserExists {
		return getErrorResponse(500, "No such user found, unable to update user data.")
	}

	//Update existing user
	updateUser := database.User{Id: userToUpdate.Id, Email: userToUpdate.Email, Name: userToUpdate.Name}
	updateUser.Update()
	if err != nil {
		return getErrorResponse(500, DEFAULT_ERROR_UPDATE_USER)
	}

	return nil
}

func RemoveUser(userId string) *models.ErrorResponse {
	userRepository = &user

	//Check if user already exists
	isUserExists, err := IsExists(userId)
	if err != nil {
		return getErrorResponse(500, "Failed to get user information. Unable to delte user.")
	}
	if !isUserExists {
		return getErrorResponse(500, "No such user found, unable to delete user.")
	}

	//Delete existing user
	err = userRepository.Remove(userId)
	if err != nil {
		return getErrorResponse(500, DEFAULT_ERROR_REMOVE_USER)
	}

	return nil
}

//TODO: Can be moved to Util
func getErrorResponse(errorCode int, errorMessage string) (*models.ErrorResponse) {
	errerResponse := models.ErrorResponse{Code: errorCode, Message: errorMessage}
	return &errerResponse;
}