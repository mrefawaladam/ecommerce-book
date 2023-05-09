package service

import (
	"ebook/internal/entity"
	"strconv"
)

func CalculateTotal(orderItems []entity.OrderItem) (int, int, error) {
	totalPrice := 0
	totalQty := 0
	for _, item := range orderItems {
		qty, err := strconv.Atoi(strconv.Itoa(int(item.Quantity)))
		if err != nil {
			return 0, 0, err
		}
		price, err := strconv.Atoi(strconv.FormatInt(item.Price, 10))
		if err != nil {
			return 0, 0, err
		}
		totalPrice += qty * price
		totalQty += qty
	}
	return totalPrice, totalQty, nil
}
