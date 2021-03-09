package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/miguelhun/go-microservices/mvc/services"
	"github.com/miguelhun/go-microservices/mvc/utils"
)

func GetUser(res http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		userError := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad request",
		}
		res.WriteHeader(userError.StatusCode)
		jsonValue, _ := json.Marshal(userError)
		res.Write(jsonValue)
		return
	}

	user, userError := services.UserService.GetUser(userId)
	if userError != nil {
		jsonValue, _ := json.Marshal(userError)
		res.WriteHeader(userError.StatusCode)
		res.Write(jsonValue)
		return
	}

	jsonValue, _ := json.Marshal(user)
	res.Write(jsonValue)
}
