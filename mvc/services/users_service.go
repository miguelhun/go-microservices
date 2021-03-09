package services

import (
	"github.com/miguelhun/go-microservices/mvc/domain"
	"github.com/miguelhun/go-microservices/mvc/utils"
)

type userService struct{}

var UserService userService

func (u *userService) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.UserDao.GetUser(userId)
}
