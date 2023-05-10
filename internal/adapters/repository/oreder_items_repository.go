package repository

import (
	"ebook/internal/entity"

	"gorm.io/gorm"
)

type OrderItemsRepository struct {
	DB *gorm.DB
}

func (repo OrderItemsRepository) CreateOrderItem(order entity.OrderItem) error {
	result := repo.DB.Create(&order)
	return result.Error
}
