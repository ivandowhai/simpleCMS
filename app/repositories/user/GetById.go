package user

import "errors"
import "../../core"
import "../../models"

func GetById(id uint64) (*models.User, error) {
	row := core.GetDB().QueryRow("select id, name, email, password, role from users where id = ?", id)

	user := new(models.User)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return &models.User{}, errors.New("user not found")
	}

	return user, nil
}
