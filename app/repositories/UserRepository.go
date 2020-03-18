package repositories

import (
	"../core"
	"../models"
	"errors"
	"fmt"
)

type UserRepository struct{}

func (r *UserRepository) GetById(id uint64) (*models.User, error) {
	row := core.GetDB().QueryRow("select id, name, email, password, role from users where id = ?", id)

	user := new(models.User)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return &models.User{}, errors.New("user not found")
	}

	return user, nil
}

func (r *UserRepository) Confirm(user *models.User) {
	_, err := core.GetDB().Exec("update users set confirmation_code = null where id = ?", user.ID)
	if err != nil {
		fmt.Println(err)
	}
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	row := core.GetDB().QueryRow("select id, name, email, password, role, confirmation_code from users where email = ?", email)

	user := new(models.User)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &user.ConfirmationCode)
	if err != nil {
		return &models.User{}, errors.New("User not found")
	}

	return user, nil
}

func (r *UserRepository) CreateUser(user models.User) {
	_, err := core.GetDB().Exec("insert into users (name, email, password, role, confirmation_code) values (?, ?, ?, ?, ?)", user.Name, user.Email, user.Password, user.Role, user.ConfirmationCode)
	if err != nil {
		fmt.Println(err)
	}
}
