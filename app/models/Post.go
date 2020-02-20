package models

import "errors"

type Post struct {
	ID      int32
	Title   string
	Content string
}

var posts = []*Post{
	{ID: 1, Title: "Test post 1", Content: "Test post 1"},
	{ID: 2, Title: "Test post 2", Content: "Test post 2"},
}

func GetAllPosts() []*Post {
	// TODO: make real
	return posts
}

func GetPostById(ID int32) (*Post, error) {
	// TODO: make real
	for i := range posts {
		if posts[i].ID == ID {
			return posts[i], nil
		}
	}

	return nil, errors.New("post not found")
}
