package post

import (
	"../../core"
	"fmt"
)

func Delete(id uint64) {
	_, err := core.GetDB().Exec("delete from posts where id = ?", id)
	if err != nil {
		fmt.Println(err)
	}
}
