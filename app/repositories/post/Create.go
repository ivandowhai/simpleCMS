package post

import (
	"../../core"
	"../../models"
	"fmt"
)

func Create(post models.Post) {
	_, err := core.GetDB().Exec("insert into posts (title, content, user_id) values (?, ?, ?)", post.Title, post.Content, post.AuthorID)
	if err != nil {
		fmt.Println(err)
	}
}
