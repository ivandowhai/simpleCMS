package repositories

import (
	"../core"
	"../models"
	"errors"
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
		r.log(err)
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
		r.log(err)
	}
}

func (r *UserRepository) GetAll() []*models.User {
	rows, err := core.GetDB().Query("select id, name, email, role from users")
	if err != nil {
		r.log(err)
	}
	defer rows.Close()

	users := make([]*models.User, 0)
	for rows.Next() {
		user := new(models.User)
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Role)
		if err != nil {
			r.log(err)
			panic(err)
		}

		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		r.log(err)
	}

	return users
}

func (r *UserRepository) ChangeRole(user *models.User) {
	_, err := core.GetDB().Exec("update users set role = ? where id = ?", user.Role, user.ID)
	if err != nil {
		r.log(err)
	}
}

func (r *UserRepository) Delete(user *models.User) {
	_, err := core.GetDB().Exec("delete from users where id = ?", user.ID)
	if err != nil {
		r.log(err)
	}
}

func (r *UserRepository) log(error error) {
	logger := core.Logger{}
	logger.Init()
	logger.WriteLog(error.Error(), "database")
}
