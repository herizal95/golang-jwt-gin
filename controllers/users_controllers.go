package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/herizal95/golang-jwt-gin/data/response"
	"github.com/herizal95/golang-jwt-gin/helper"
	"github.com/herizal95/golang-jwt-gin/models"
	"github.com/herizal95/golang-jwt-gin/repository"
)

type UserController struct {
	userRepository repository.UserRepository
}

func NewUserController(repository repository.UserRepository) *UserController {
	return &UserController{userRepository: repository}
}

// Get all users
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

// Get user by Id
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

// Create new users
func (controller *UserController) CreateUser(ctx *gin.Context) {
	var users models.Users
	err := ctx.ShouldBindJSON(&users)
	helper.ErrorPanic(err)

	controller.userRepository.Save(users)

	webResponse := response.Response{
		Code:    200,
		Status:  "Ok",
		Message: "Successfully created users!",
		Data:    nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

// Delete Users by Id
func (controller *UserController) DeleteUserId(ctx *gin.Context) {

	user_id := ctx.Param("id")

	rowsAffected, err := controller.userRepository.Delete(user_id)

	if err != nil {
		webResponse := response.Response{
			Code:    http.StatusInternalServerError,
			Status:  "Error",
			Message: "Failed to delete user",
			Data:    nil,
		}
		ctx.JSON(http.StatusInternalServerError, webResponse)
		return
	}

	webResponse := response.Response{
		Code:    http.StatusOK,
		Status:  "Ok!",
		Message: "User has been deleted!",
		Data:    rowsAffected,
	}

	ctx.JSON(http.StatusOK, webResponse)
}
