package postgres

import (
	database "lifesaver/internal/database/postgres"
	"lifesaver/pkg/models"
)

type User models.User
type UserRepository models.UserRepository

const (
	IS_EXISTS		= "SELECT id FROM lifesaver.user WHERE id = $1 or email = $2 LIMIT 1"
	CREATE_USER		= "INSERT INTO lifesaver.user(id, emal, name) VALUES ($1, $2, $3) RETURNING id"
	GET_USER		= "SELECT id, email, name FROM lifesaver.user WHERE id = $1 LIMIT 1"
	UPDATE_USER 	= "UPDATE lifesaver.user SET email = $1, name = $2 WHERE id = $3"
	DELETE_USER		= "DELTE FROM lifesaver.user WHERE id = $1"
)

func(user User) IsExists(userId string, emailId string) (bool, error) {
	err := database.Client.QueryRow(IS_EXISTS, userId, emailId).Scan(&userId)
	
	if err != nil {
		return false, err
	}

	if len(userId) > 0 {
		return false, nil
	}

	return true, nil
}

func (user *User) Save() (string, error) {
	var userId string
	err := database.Client.QueryRow(CREATE_USER, user.Id, user.Email, user.Name).Scan(&userId)
	
	if err != nil {
		return "", err
	}

	return userId, err
}

func(user User) GetById(userId string) (*models.User, error) {
	err := database.Client.QueryRow(GET_USER, userId).Scan(&user.Id, &user.Email, &user.Name)
	
	if err != nil {
		return nil, err
	}

	userReturn := models.User(user)

	return &userReturn, nil
}

func(user *User) Update() (error) {
	_, err := database.Client.Exec(UPDATE_USER, user.Email, user.Name, user.Id)
	
	if err != nil {
		return err
	}

	return nil
}

func(user User) Remove(userId string) (error) {
	_, err := database.Client.Exec(DELETE_USER, userId)
	
	if err != nil {
		return err
	}

	return nil
}