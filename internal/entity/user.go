package entity

import (
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model

	Name     string    `json:"name" form:"name" validate:"required"`
	Email    string    `json:"email" form:"email" validate:"required,email" `
	Password string    `json:"password" form:"password" validate:"required"`
	Role     string    `gorm:"default:customer"`
	Address  []Address `gorm:"foreignKey:UserID"`
}
