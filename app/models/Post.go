package models

import (
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)
import "../core"

type Post struct {
	ID      uint64
	Title   string
	Content string
}

var posts = []*Post{
	{ID: 1, Title: "Test post 1", Content: "Test post 1"},
	{ID: 2, Title: "Test post 2", Content: "Test post 2"},
}

func GetAllPosts() []*Post {
	rows, err := core.GetDB().Query("select id, title, content from posts")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	posts := make([]*Post, 0)
	for rows.Next() {
		post := new(Post)
		err := rows.Scan(&post.ID, &post.Title, &post.Content)
		if err != nil {
			panic(err)
		}

		posts = append(posts, post)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err)
	}

	return posts
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
