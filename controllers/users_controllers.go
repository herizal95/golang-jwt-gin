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

func (controller *UserController) GetUsersById(ctx *gin.Context) {
	user_id := ctx.Param("id")
	users, err := controller.userRepository.FindById(user_id)

	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusNotFound,
			Status:  "Error",
			Message: "Id User not found!",
			Data:    nil,
		}
		ctx.JSON(http.StatusNotFound, webResponse)
		return

	} else {
		webResponse := response.Response{
			Code:    200,
			Status:  "Ok",
			Message: "Successfully fetch user data!",
			Data:    users,
		}

		ctx.JSON(http.StatusOK, webResponse)

	}
}
