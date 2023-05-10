package entity

import (
	"gorm.io/gorm"
)

type Payment struct {
	*gorm.Model

	PaymentToken  string `json:"payment_token" form:"payment_token" `
	PaymentType   string `json:"payment_type" form:"payment_type" `
	PaymentDate   string `json:"payment_date" form:"payment_date" `
	PaymentStatus string `json:"payment_status" form:"payment_status" `
	PaymentExpiry string `json:"payment_expiry" form:"payment_expiry" `
	OrderId       int    `json:"order_id" form:"order_id" `
}
