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

func (r *UserRepository) GetAll() []*models.User {
	rows, err := core.GetDB().Query("select id, name, email, role from users")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	users := make([]*models.User, 0)
	for rows.Next() {
		user := new(models.User)
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Role)
		if err != nil {
			panic(err)
		}

		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err)
	}

	return users
}
