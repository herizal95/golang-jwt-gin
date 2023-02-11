package service

import "github.com/herizal95/golang-jwt-gin/data/request"

type AuthenticationService interface {
	Login(users request.LoginRequest) (string, error)
	Register(users request.CreateUserRequest)
}
