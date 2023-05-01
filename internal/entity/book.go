package entity

import (
	"gorm.io/gorm"
)

type Book struct {
	*gorm.Model

	Title       string   `json:"title" form:"title" `
	Author      string   `json:"author" form:"author" `
	Cover       string   `json:"cover" form:"cover" `
	Status      string   `json:"status" form:"status" `
	Description string   `json:"description" form:"description" `
	Stok        string   `json:"stok" form:"stok" `
	Price       string   `json:"price" form:"price" `
	Weight      string   `json:"weight" form:"weight" `
	CategoryID  uint     `json:"category_id" form:"category_id" `
	SellerId    uint     `json:"seller_id" form:"seller_id" `
	Category    Category `json:"category" gorm:"foreignKey:CategoryID"`
}
