package entity

import (
	"gorm.io/gorm"
)

type OrderItem struct {
	*gorm.Model

	OrderId  uint   `json:"order_id" form:"order_id" `
	BookId   string `json:"book_id" form:"book_id" `
	Quantity int32  `json:"quantity" form:"quantity"`
	Price    int64  `json:"price" form:"price" `
	Book     Book   `json:"book" gorm:"foreignKey:BookId"`
}
