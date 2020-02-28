package user

import (
	"../../core"
	"../../models"
	"errors"
)

func GetByEmailAndPassword(email string, password string) (*models.User, error) {
	row := core.GetDB().QueryRow("select id, name, email, role from users where email = :email and password = :password", email, password)

	user := new(models.User)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Role)
	if err != nil {
		return &models.User{}, errors.New("user not found")
	}

	return user, nil
}
