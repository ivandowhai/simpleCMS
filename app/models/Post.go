package models

import _ "github.com/go-sql-driver/mysql"

type Post struct {
	ID       uint64 `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorID uint64 `json:"authorId"`
}
