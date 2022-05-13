package postgres

import (
	database "lifesaver/internal/database/postgres"
	"lifesaver/pkg/models"
)

type User models.User
type UserRepository models.UserRepository

const (
	IS_EXISTS		= "SELECT id FROM users WHERE id = $1 LIMIT 1"
	CREATE_USER		= "INSERT INTO users(id, email, name) VALUES ($1, $2, $3) RETURNING id"
	GET_USER		= "SELECT id, email, name FROM users WHERE id = $1 LIMIT 1"
	UPDATE_USER 	= "UPDATE users SET email = $1, name = $2 WHERE id = $3"
	DELETE_USER		= "DELETE FROM users WHERE id = $1"
)

func(user User) IsExists(userId string) (bool, error) {
	rows, err := database.Client.Query(IS_EXISTS, userId)
	
	//ERROR
	if err != nil {
		return false, err
	}

	var fetchedUserId string
	for rows.Next() {
		err = rows.Scan(&fetchedUserId)

		//ERROR
		if err != nil {
			return false, nil
		}
	}

	//FOUND
	if len(fetchedUserId) > 0 {
		return true, nil
	}

	//NOT FOUND
	return false, nil
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
	rows, err := database.Client.Query(GET_USER, userId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Email, &user.Name)
		
		//ERROR
		if err != nil {
			return nil, nil
		}
	}

	//NOT FOUND
	if user.Id == "" {
		return nil, nil
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