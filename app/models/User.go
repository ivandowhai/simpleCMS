package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID               uint64
	Name             string
	Email            string
	Password         string
	Role             uint8
	ConfirmationCode sql.NullString
}
