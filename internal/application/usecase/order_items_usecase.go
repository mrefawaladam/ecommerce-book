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
