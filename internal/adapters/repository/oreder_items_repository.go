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

func (repo OrderItemsRepository) GetOrderItemsByBook(id int) ([]entity.OrderItem, error) {
	var orderItems []entity.OrderItem
	result := repo.DB.Where("book_id = ?", id).Find(&orderItems)
	return orderItems, result.Error
}
