package models

import (
	"gorm.io/gorm"
)

type UserModel struct {
	*gorm.Model

	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required"`
}
