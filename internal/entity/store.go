package entity

import (
	"gorm.io/gorm"
)

type Store struct {
	*gorm.Model

	Name        string `json:"name" form:"total_price" `
	Address     string `json:"address" form:"address" `
	Profile     string `json:"profile" form:"profile" `
	PhoneNumber string `json:"phone_number" form:"phone_number" `
	SellerId    uint   `json:"seller_id" form:"seller_id" `
}
