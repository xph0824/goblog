package models

import (
	"goModWork/databases"
)


// Create create user on form
func CreateUser(user databases.User) error {
	return databases.UserDAO.CreateUser(databases.User{
		Name: user.Name,
		Age: user.Age,
		Address: user.Address,
		Email: user.Email,
		Phone: user.Phone,
	})
}

// First get first user from database
func FirstUser() (databases.User, error) {
	return databases.UserDAO.FirstUser()
}

// Update update user records
func UpdateUser(user databases.User) error {
	return databases.UserDAO.UpdateUser(databases.User{
		Name: user.Name,
		Age: user.Age,
		Address: user.Address,
		Email: user.Email,
		Phone: user.Phone,
	})
}

// Delete delete user records
func DeleteUser() error {
	return databases.UserDAO.DeleteUser()
}
