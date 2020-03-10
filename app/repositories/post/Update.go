package post

import (
	"../../core"
	"fmt"
)

func Update(id uint64, title string, content string) {
	_, err := core.GetDB().Exec("update posts set title = ?, content = ? where id = ?", title, content, id)
	if err != nil {
		fmt.Println(err)
	}
}
