package entity

import (
	"gorm.io/gorm"
)

type Order struct {
	*gorm.Model

	TotalPrice      string `json:"total_price" form:"total_price" `
	PaymentExpiry   string `json:"payment_expiry" form:"payment_expiry" `
	ShippingAddress string `json:"shipping_address" form:"shipping_address" `
	CustomerID      uint   `json:"customer_id" form:"customer_id" `
}
