package user

import (
	"../../core"
	"../../models"
	"errors"
)

func GetByEmail(email string) (*models.User, error) {
	row := core.GetDB().QueryRow("select id, name, email, password, role, confirmation_code from users where email = ?", email)

	user := new(models.User)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role, &user.ConfirmationCode)
	if err != nil {
		return &models.User{}, errors.New("User not found")
	}

	return user, nil
}
