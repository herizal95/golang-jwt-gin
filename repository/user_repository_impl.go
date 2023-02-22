package repository

import (
	"errors"

	"github.com/google/uuid"
	"github.com/herizal95/golang-jwt-gin/data/request"
	"github.com/herizal95/golang-jwt-gin/helper"
	"github.com/herizal95/golang-jwt-gin/models"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserRepositoryImpl(Db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{Db: Db}
}

// Delete implements UserRepository
func (ctx *UserRepositoryImpl) Delete(usersId uuid.UUID) {
	var users models.Users
	result := ctx.Db.Where("id = ?", usersId).Delete(&users)
	helper.ErrorPanic(result.Error)
}

// FindAll implements UserRepository
func (ctx *UserRepositoryImpl) FindAll() []models.Users {
	var users []models.Users
	result := ctx.Db.Find(&users)
	helper.ErrorPanic(result.Error)
	return users
}

// FindById implements UserRepository
func (ctx *UserRepositoryImpl) FindById(usersId string) (models.Users, error) {
	var users models.Users
	result := ctx.Db.Raw("SELECT * FROM users WHERE id = ?", usersId).Scan(&users)
	if result.Error != nil {
		return users, result.Error
	}

	return users, nil
}

// FindByUsername implements UserRepository
func (ctx *UserRepositoryImpl) FindByUsername(username string) (models.Users, error) {
	var users models.Users
	result := ctx.Db.First(&users, "username = ?", username)

	if result.Error != nil {
		return users, errors.New("invalid username or password")
	}
	return users, nil
}

// Save implements UserRepository
func (ctx *UserRepositoryImpl) Save(users models.Users) {
	data := models.Users{
		Id:       uuid.New(),
		Username: users.Username,
		Email:    users.Email,
		Password: users.Password,
	}
	result := ctx.Db.Create(&data)
	helper.ErrorPanic(result.Error)
}

// Update implements UserRepository
func (ctx *UserRepositoryImpl) Update(users models.Users) {
	var UpdateUsers = request.UpdateUserRequest{
		Id:       users.Id,
		Username: users.Username,
		Email:    users.Email,
		Password: users.Password,
	}
	result := ctx.Db.Model(&users).Updates(&UpdateUsers)
	helper.ErrorPanic(result.Error)
}
