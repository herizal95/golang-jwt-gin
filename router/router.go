package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/herizal95/golang-jwt-gin/controllers"
	"github.com/herizal95/golang-jwt-gin/middleware"
	"github.com/herizal95/golang-jwt-gin/repository"
)

func NewRouter(userRepository repository.UserRepository, authControllers *controllers.AuthenticationController, userController *controllers.UserController) *gin.Engine {
	service := gin.Default()

	service.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome to Golang")
	})

	router := service.Group("/api")

	authenticationRouter := router.Group("/auth")
	authenticationRouter.POST("/register", authControllers.Register)
	authenticationRouter.POST("/login", authControllers.Login)

	usersRouter := router.Group("/users")
	usersRouter.GET("", middleware.DeserializerUser(userRepository), userController.GetUsers)
	usersRouter.GET("/:id", middleware.DeserializerUser(userRepository), userController.GetUsersById)

	return service
}
