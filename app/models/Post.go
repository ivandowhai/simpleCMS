package models

import _ "github.com/go-sql-driver/mysql"

type Post struct {
	ID      uint64
	Title   string
	Content string
}
