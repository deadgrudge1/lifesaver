package models

import "strings"

type User struct {
	Id 		string		`json:"id"`
	Email 	string		`json:"email"`
	Name 	string		`json:"name"`
}

type UserRepository interface {

	//Get existing user details
	IsExists(userId string, email string) (bool, error)

	//Create a new user
	Save() (string, error)

	//Get existing user details
	GetById(userId string) (*User, error)

	//Update details for existing user
	Update() (error)

	//Remove existing user
	Remove(userId string) (error)

}

func (user *User) ValidateInsert() string {
	var errorMessage string

	if len(strings.TrimSpace(user.Email)) < 1 {
		appendToError(&errorMessage, "User Email is required.")
	}

	if len(strings.TrimSpace(user.Name)) < 1 {
		appendToError(&errorMessage, "User Name is required.")
	}
	
	return errorMessage;
}

//TODO: Can be moved to Util
func appendToError(errorMessage *string, messageToAppend string) {
	if len(*errorMessage) > 1 {
		*errorMessage += " "
	}

	*errorMessage += messageToAppend
}