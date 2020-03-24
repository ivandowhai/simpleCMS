package services

import (
	"../core"
	"../models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// TODO: real secret
var jwtKey = []byte("my_secret_key")

type Claims struct {
	ID              uint64 `json:"id"`
	UserRole        uint8  `json:"userRole"`
	IsUserConfirmed bool   `json:"isUserConfirmed"`
	jwt.StandardClaims
}

type LoginService struct {
	ExpirationTime time.Time
	TokenString    string
}

func (ls *LoginService) Login(user *models.User) error {
	logger := core.Logger{}
	logger.Init()

	ls.ExpirationTime = time.Now().Add(5 * time.Minute)
	claims := &Claims{
		ID:              user.ID,
		UserRole:        user.Role,
		IsUserConfirmed: !user.ConfirmationCode.Valid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: ls.ExpirationTime.Unix(),
		}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		logger.WriteLog(err.Error(), "error")
		return err
	}

	ls.TokenString = tokenString

	return nil
}
