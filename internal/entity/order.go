package entity

import (
	"gorm.io/gorm"
)

type Order struct {
	*gorm.Model

	TotalOrder      string      `json:"total_order" form:"total_order" `
	ShippingAddress string      `json:"shipping_address" form:"shipping_address"`
	TotalShipping   string      `json:"total_shipping" form:"total_shipping" `
	CustomerID      uint        `json:"customer_id" form:"customer_id" `
	OrderItems      []OrderItem `json:"order_items" gorm:"foreignKey:OrderId"`
}
