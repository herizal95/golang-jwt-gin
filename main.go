package main

import (
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/herizal95/golang-jwt-gin/config"
	"github.com/herizal95/golang-jwt-gin/controllers"
	"github.com/herizal95/golang-jwt-gin/helper"
	"github.com/herizal95/golang-jwt-gin/models"
	"github.com/herizal95/golang-jwt-gin/repository"
	"github.com/herizal95/golang-jwt-gin/router"
	"github.com/herizal95/golang-jwt-gin/service"
)

func main() {

	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load environment variable", err)
	}

	// database
	db := config.ConnectionDB(&loadConfig)
	validate := validator.New()

	db.Table("users").AutoMigrate(&models.Users{})

	// init repository
	usersRepository := repository.NewUserRepositoryImpl(db)

	// Init serveice
	authenticationService := service.NewAuthenticationLogin(usersRepository, validate)

	// Init Controller
	authenticationControllers := controllers.NewAuthenticationController(authenticationService)
	userController := controllers.NewUserController(usersRepository)

	routes := router.NewRouter(usersRepository, authenticationControllers, userController)

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}

	server_err := server.ListenAndServe()
	helper.ErrorPanic(server_err)
}
