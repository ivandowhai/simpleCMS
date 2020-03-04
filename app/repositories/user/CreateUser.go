package user

import (
	"../../core"
	"../../models"
	"fmt"
)

func CreateUser(user models.User) {
	_, err := core.GetDB().Exec("insert into users set name = :name, email = :email, password := password, role = :role", user)
	if err != nil {
		fmt.Println(err)
	}
}
