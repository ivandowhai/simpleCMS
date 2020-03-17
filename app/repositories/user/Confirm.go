package user

import (
	"../../core"
	"../../models"
	"fmt"
)

func Confirm(user *models.User) {
	_, err := core.GetDB().Exec("update users set confirmation_code = null where id = ?", user.ID)
	if err != nil {
		fmt.Println(err)
	}
}
