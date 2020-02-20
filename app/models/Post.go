package models

type Post struct {
	Title   string
	Content string
}

func GetAllPosts() []*Post {
	var posts = []*Post{
		{Title: "Test post 1", Content: "Test post 1"},
		{Title: "Test post 2", Content: "Test post 2"},
	}
	return posts
}
