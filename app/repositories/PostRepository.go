package repositories

import (
	"../core"
	"../models"
	"errors"
	"fmt"
)

type PostRepository struct{}

func (r *PostRepository) Create(post models.Post) {
	_, err := core.GetDB().Exec("insert into posts (title, content, user_id) values (?, ?, ?)", post.Title, post.Content, post.AuthorID)
	if err != nil {
		fmt.Println(err)
	}
}

func (r *PostRepository) Delete(id uint64) {
	_, err := core.GetDB().Exec("delete from posts where id = ?", id)
	if err != nil {
		fmt.Println(err)
	}
}

func (r *PostRepository) GetAll() []*models.Post {
	rows, err := core.GetDB().Query("select id, title, content from posts")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	posts := make([]*models.Post, 0)
	for rows.Next() {
		post := new(models.Post)
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

func (r *PostRepository) GetByUser(userID uint64) []*models.Post {
	rows, err := core.GetDB().Query("select id, title from posts where user_id = ?", userID)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	posts := make([]*models.Post, 0)
	for rows.Next() {
		post := new(models.Post)
		err := rows.Scan(&post.ID, &post.Title)
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

func (r *PostRepository) GetOne(id uint64) (*models.Post, error) {
	row := core.GetDB().QueryRow("select id, title, content from posts where id = ?", id)

	post := new(models.Post)
	err := row.Scan(&post.ID, &post.Title, &post.Content)
	if err != nil {
		return post, errors.New("post not found")
	}

	return post, nil
}

func (r *PostRepository) Update(id uint64, title string, content string) {
	_, err := core.GetDB().Exec("update posts set title = ?, content = ? where id = ?", title, content, id)
	if err != nil {
		fmt.Println(err)
	}
}
