package domain

import (
	"fmt"
	"net/http"

	"github.com/miguelhun/go-microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "Miguel", LastName: "Hun", Email: "email@gmail.com"},
	}

	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct{}

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	user := users[userId]
	if user == nil {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("user %v was not found", userId),
			StatusCode: http.StatusNotFound,
			Code:       "not found",
		}
	}

	return user, nil
}
