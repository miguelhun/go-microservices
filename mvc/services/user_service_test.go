package services

import (
	"net/http"
	"testing"

	"github.com/miguelhun/go-microservices/mvc/domain"
	"github.com/miguelhun/go-microservices/mvc/utils"
	"github.com/stretchr/testify/assert"
)

var (
	userDaoMock usersDaoMock

	getUserFunction func(userId int64) (*domain.User, *utils.ApplicationError)
)

func init() {
	domain.UserDao = &usersDaoMock{}
}

type usersDaoMock struct{}

func (m *usersDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userId)
}

func TestGetUserNotFound(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			StatusCode: http.StatusNotFound,
			Message:    "user 0 was not found",
		}
	}
	user, err := UserService.GetUser(0)

	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 was not found", err.Message)

}

func TestGetUserNoError(t *testing.T) {
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
			Id:        123,
			FirstName: "Miguel",
			LastName:  "Hun",
			Email:     "email@gmail.com",
		}, nil
	}
	user, err := UserService.GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Miguel", user.FirstName)
	assert.EqualValues(t, "Hun", user.LastName)
	assert.EqualValues(t, "email@gmail.com", user.Email)
}
