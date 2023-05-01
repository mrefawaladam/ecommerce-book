package entity

import (
	"gorm.io/gorm"
)

type Address struct {
	*gorm.Model

	Street     string `json:"street" form:"street" `
	City       string `json:"city" form:"city" `
	Province   string `json:"province" form:"province" `
	Country    string `json:"country" form:"country" `
	PostalCode string `json:"postal_code" form:"postal_code" `
	UserID     uint   `json:"user_id" form:"user_id"`
}
