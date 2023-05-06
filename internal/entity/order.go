package entity

import (
	"gorm.io/gorm"
)

type Order struct {
	*gorm.Model

	TotalOrder      string `json:"total_order" form:"total_order" `
	PaymentExpiry   string `json:"payment_expiry" form:"payment_expiry" `
	ShippingAddress string `json:"shipping_address" form:"shipping_address"`
	TotalShipping   string `json:"total_shipping" form:"total_shipping" `
	NoResi          string `json:"no_resi" form:"no_resi" `
	CustomerID      uint   `json:"customer_id" form:"customer_id" `
}
