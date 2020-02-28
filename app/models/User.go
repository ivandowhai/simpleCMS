package models

import _ "github.com/go-sql-driver/mysql"

type User struct {
	ID       uint64
	Name     string
	Email    string
	Password string
	Role     uint8
}
