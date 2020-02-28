package post

import (
	"../../core"
	"../../models"
	"errors"
)

func GetOne(id uint64) (*models.Post, error) {
	row := core.GetDB().QueryRow("select id, title, content from posts where id = ?", id)

	post := new(models.Post)
	err := row.Scan(&post.ID, &post.Title, &post.Content)
	if err != nil {
		return post, errors.New("post not found")
	}

	return post, nil
}
