package entity

import (
	"errors"

	"gorm.io/gorm"
)

type Book struct {
	*gorm.Model

	Title       string   `json:"title" form:"title validate:"required"`
	Author      string   `json:"author" form:"author validate:"required"`
	Cover       string   `json:"cover" form:"cover validate:"required"`
	Status      string   `json:"status" form:"status" validate:"required"`
	Description string   `json:"description" form:"description" validate:"required"`
	Stok        uint     `json:"stok" form:"stok" validate:"required"`
	Price       float64  `json:"price" form:"price" validate:"required"`
	Weight      float64  `json:"weight" form:"weight" validate:"required"`
	CategoryID  uint     `json:"category_id" form:"category_id" validate:"required"`
	SellerId    uint     `json:"seller_id" form:"seller_id"`
	Category    Category `json:"category" gorm:"foreignKey:CategoryID"`
}

func (b *Book) DecreaseStock(quantity uint) error {
	if b.Stok < quantity {
		return errors.New("not enough stock")
	}
	b.Stok -= quantity
	return nil
}

func (b *Book) IncreaseStock(quantity uint) {
	b.Stok += quantity
}
