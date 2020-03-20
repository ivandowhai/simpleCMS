package services

import (
	"../core"
	"../repositories"
	"strconv"
)

type UserService struct{}

func (s *UserService) ChangeRole(userIdStr string, roleStr string) {
	logger := core.Logger{}
	logger.Init()
	userRepository := repositories.UserRepository{}

	ID, err := strconv.ParseUint(userIdStr, 10, 16)
	if err != nil {
		logger.WriteLog(err.Error(), "error")
		return
	}

	user, err := userRepository.GetById(ID)
	if err != nil {
		logger.WriteLog(err.Error(), "error")
		return
	}

	role, err := strconv.ParseUint(roleStr, 10, 16)
	if err != nil {
		logger.WriteLog(err.Error(), "error")
		return
	}

	user.Role = uint8(role)

	userRepository.ChangeRole(user)
}

func (s *UserService) Delete(userIdStr string) {
	logger := core.Logger{}
	logger.Init()
	userRepository := repositories.UserRepository{}
	ID, err := strconv.ParseUint(userIdStr, 10, 16)
	if err != nil {
		logger.WriteLog(err.Error(), "error")
		return
	}

	user, err := userRepository.GetById(ID)
	if err != nil {
		logger.WriteLog(err.Error(), "error")
		return
	}

	userRepository.Delete(user)
}
