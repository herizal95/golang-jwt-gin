package repository

import (
	"github.com/google/uuid"
	"github.com/herizal95/golang-jwt-gin/models"
)

type UserRepository interface {
	Save(users models.Users)
	Update(users models.Users)
	Delete(usersId uuid.UUID)
	FindById(usersId uuid.UUID) (models.Users, error)
	FindAll() []models.Users
	FindByUsername(username string) (models.Users, error)
}
