package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/miguelhun/go-microservices/mvc/services"
	"github.com/miguelhun/go-microservices/mvc/utils"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		userError := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad request",
		}
		utils.RespondError(c, userError)
		return
	}

	user, userError := services.UserService.GetUser(userId)
	if userError != nil {
		utils.RespondError(c, userError)
		return
	}

	utils.Respond(c, http.StatusOK, user)
}
