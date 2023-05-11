package usecase

import (
	"ebook/internal/adapters/repository"
	"ebook/internal/entity"
)

type OrderItemsUsecase struct {
	Repo repository.OrderItemsRepository
}

func (usecase OrderItemsUsecase) CreateOrderItems(order entity.OrderItem) error {
	err := usecase.Repo.CreateOrderItem(order)
	return err
}
func (usecase OrderItemsUsecase) GetOrderItemByBook(id int) ([]entity.OrderItem, error) {
	orderItems, err := usecase.Repo.GetOrderItemsByBook(id)
	if err != nil {
		return nil, err
	}
	return orderItems, nil
}

func (u *OrderItemsUsecase) GetBestSellingBooks() ([]entity.OrderItem, error) {
	var orderItems []entity.OrderItem
	err := u.Repo.DB.
		Model(&entity.OrderItem{}).
		Select("book_id, SUM(quantity) as total_quantity").
		Group("book_id").
		Order("total_quantity DESC").
		Limit(10).
		Find(&orderItems).
		Error
	if err != nil {
		return nil, err
	}

	return orderItems, nil
}
