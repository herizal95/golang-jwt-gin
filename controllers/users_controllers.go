package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/herizal95/golang-jwt-gin/data/response"
	"github.com/herizal95/golang-jwt-gin/repository"
)

type UserController struct {
	userRepository repository.UserRepository
}

func NewUserController(repository repository.UserRepository) *UserController {
	return &UserController{userRepository: repository}
}

func (controllers *UserController) GetUsers(ctx *gin.Context) {
	users := controllers.userRepository.FindAll()
	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfuly fetch all data user!",
		Data:    users,
	}

	ctx.JSON(http.StatusOK, webResponse)

}
