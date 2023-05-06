package entity

import (
	"gorm.io/gorm"
)

type OrderItem struct {
	*gorm.Model

	OrderId  string `json:"order_id" form:"order_id" `
	BookId   string `json:"book_id" form:"book_id" `
	Quantity string `json:"quantity" form:"quantity"`
	Price    string `json:"price" form:"price" `
}
