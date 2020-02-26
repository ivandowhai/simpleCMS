package models

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

type Post struct {
	ID      uint64
	Title   string
	Content string
}

var posts = []*Post{
	{ID: 1, Title: "Test post 1", Content: "Test post 1"},
	{ID: 2, Title: "Test post 2", Content: "Test post 2"},
}

func GetPostById(ID uint64) (*Post, error) {
	// TODO: make real
	for i := range posts {
		if posts[i].ID == ID {
			return posts[i], nil
		}
	}

	return nil, errors.New("post not found")
}
