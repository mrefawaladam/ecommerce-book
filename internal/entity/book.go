package entity

import (
	"gorm.io/gorm"
)

type Book struct {
	*gorm.Model

	Title       string   `json:"title" form:"title validate:"required"`
	Author      string   `json:"author" form:"author validate:"required"`
	Cover       string   `json:"cover" form:"cover validate:"required" `
	Status      string   `json:"status" form:"status" validate:"required" `
	Description string   `json:"description" form:"description" validate:"required"`
	Stok        string   `json:"stok" form:"stok" validate:"required"`
	Price       string   `json:"price" form:"price" validate:"required"`
	Weight      string   `json:"weight" form:"weight" validate:"required"`
	CategoryID  uint     `json:"category_id" form:"category_id" validate:"required"`
	SellerId    uint     `json:"seller_id" form:"seller_id"`
	Category    Category `json:"category" gorm:"foreignKey:CategoryID"`
}
