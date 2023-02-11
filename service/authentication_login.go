package service

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/herizal95/golang-jwt-gin/config"
	"github.com/herizal95/golang-jwt-gin/data/request"
	"github.com/herizal95/golang-jwt-gin/helper"
	"github.com/herizal95/golang-jwt-gin/models"
	"github.com/herizal95/golang-jwt-gin/repository"
	"github.com/herizal95/golang-jwt-gin/utils"
)

type AuthenticationLogin struct {
	UsersRepository repository.UserRepository
	Validate        *validator.Validate
}

func NewAuthenticationLogin(usersRepository repository.UserRepository, validate *validator.Validate) AuthenticationService {
	return &AuthenticationLogin{
		UsersRepository: usersRepository,
		Validate:        validate,
	}
}

// Login implements AuthenticationService
func (a *AuthenticationLogin) Login(users request.LoginRequest) (string, error) {

	// Find username in database
	new_user, user_err := a.UsersRepository.FindByUsername(users.Username)
	if user_err != nil {
		return "", errors.New("invalid username or password")
	}

	config, _ := config.LoadConfig(".")

	verity_error := utils.VerifyPassword(new_user.Password, users.Password)
	if verity_error != nil {
		return "", errors.New("invalid username or password")
	}

	//Generate token
	token, err_token := utils.GenerateToken(config.TokenExpiredIn, new_user.Id, config.TokenSecret)
	helper.ErrorPanic(err_token)
	return token, nil
}

// Register implements AuthenticationService
func (a *AuthenticationLogin) Register(users request.CreateUserRequest) {

	hashedPassword, err := utils.HashPassword(users.Password)
	helper.ErrorPanic(err)

	newUsers := models.Users{
		Username: users.Username,
		Email:    users.Email,
		Password: hashedPassword,
	}
	a.UsersRepository.Save(newUsers)
}
