package request

import "github.com/google/uuid"

type CreateUserRequest struct {
	Username string `validate:"required,min=2,max=100" json:"username"`
	Email    string `validate:"required,min=2,max=100" json:"email"`
	Password string `validate:"required,min=2,max=100" json:"password"`
}

type UpdateUserRequest struct {
	Id       uuid.UUID `validate:"required"`
	Username string    `validate:"required,min=2,max=100" json:"username"`
	Email    string    `validate:"required,min=2,max=100" json:"email"`
	Password string    `validate:"required,min=2,max=100" json:"password"`
}

type LoginRequest struct {
	Username string `validate:"required,min=2,max=100" json:"username"`
	Password string `validate:"required,min=2,max=100" json:"password"`
}
