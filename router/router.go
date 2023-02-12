package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/herizal95/golang-jwt-gin/controllers"
)

func NewRouter(authControllers *controllers.AuthenticationController) *gin.Engine {
	service := gin.Default()

	service.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome to Golang")
	})

	router := service.Group("/api")

	authRouter := router.Group("/auth")
	authRouter.POST("/register", authControllers.Register)
	authRouter.POST("/login", authControllers.Login)

	return service
}
