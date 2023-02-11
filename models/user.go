package models

import (
	"github.com/google/uuid"
)

type Users struct {
	Id       uuid.UUID `gorm:"primary_key" json:"id"`
	Username string    `gorm:"type:varchar(255);not null"`
	Email    string    `gorm:"uniqueIndex;not nul"`
	Password string    `gorm:"not null"`
}
