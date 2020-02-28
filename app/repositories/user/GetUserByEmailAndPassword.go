package user

import (
	"../../core"
	"../../models"
	"errors"
	"fmt"
)

func GetByEmailAndPassword(email string, password string) (*models.User, error) {
	rows, err := core.GetDB().Query("select id, name, email, role from users where email = \"" + email + "\" and password = \"" + password + "\"")

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

	if len(users) == 0 {
		return &models.User{}, errors.New("user not found")
	}

	return users[0], nil
}
