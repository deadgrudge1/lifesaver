package postgres

import (
	database "lifesaver/internal/database/postgres"
	"lifesaver/pkg/models"
	// "log"
)

type User models.User
type UserRepository models.UserRepository

const (
	IS_EXISTS		= "SELECT id FROM users WHERE id = $1 or email = $2 LIMIT 1"
	CREATE_USER		= "INSERT INTO users(id, email, name) VALUES ($1, $2, $3) RETURNING id"
	GET_USER		= "SELECT id, email, name FROM users WHERE id = $1 LIMIT 1"
	UPDATE_USER 	= "UPDATE users SET email = $1, name = $2 WHERE id = $3"
	DELETE_USER		= "DELETE FROM users WHERE id = $1"
)

//b57d9b6d-f3a2-4914-b5e5-ca2e6cce12b0

// func init() {
// 	_, err := database.Client.Exec("CREATE TABLE users( id varchar(64) primary key, email varchar(64), name varchar(64) )") 

// 	if err != nil {
// 		log.Println("Failed to create users table")
// 	} else {
// 		log.Println("Created users table")
// 	}
// }

func(user User) IsExists(userId string, emailId string) (bool, error) {
	var err error
	err = database.Client.QueryRow(IS_EXISTS, userId, emailId).Scan(&userId)
	
	if err != nil {
		return false, err
	}

	if len(userId) > 0 {
		return true, nil
	}

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