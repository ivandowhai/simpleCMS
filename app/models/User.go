package models

type User struct {
	ID       uint64
	Name     string
	Email    string
	Password string
	Role     uint8
}
