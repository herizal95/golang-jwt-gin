package response

import "github.com/google/uuid"

type UserResonse struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

type LoginResponse struct {
	TokenType string `json:"token_type"`
	Token     string `json:"token"`
}
