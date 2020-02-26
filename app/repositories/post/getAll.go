package post

import "fmt"
import "../../models"
import "../../core"

func GetAll() []*models.Post {
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
