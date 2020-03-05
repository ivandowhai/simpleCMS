package user

import (
	"../../core"
	"../../models"
	"fmt"
)

func CreateUser(user models.User) {
	_, err := core.GetDB().Exec("insert into users (name, email, password, role) values (?, ?, ?, ?)", user.Name, user.Email, user.Password, user.Role)
	if err != nil {
		fmt.Println(err)
	}
}
