package repository

import (
	"github.com/herizal95/golang-jwt-gin/models"
)

type UserRepository interface {
	Save(users models.Users)
	Update(users models.Users)
	Delete(usersId string) (models.Users, error)
	FindById(usersId string) (models.Users, error)
	FindAll() []models.Users
	FindByUsername(username string) (models.Users, error)
}
