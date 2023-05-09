package repository

import (
	"ebook/internal/entity"
	"errors"
	"fmt"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func (repo *OrderRepository) GetLastOrderID() (uint, error) {
	fmt.Println("repo")

	var order entity.Order
	if err := repo.DB.Order("id desc").First(&order).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 1, nil
		}
		return 0, err
	}

	return order.ID, nil
}
func (r *OrderRepository) FindByID(id uint) (*entity.Order, error) {
	var order entity.Order
	err := r.DB.First(&order, id).Error
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepository) GetUserByID(id uint) (*entity.User, error) {
	var user entity.User
	err := r.DB.Preload("Address").First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *OrderRepository) GetAddressByID(id uint) (*entity.Address, error) {
	var address entity.Address
	err := r.DB.First(&address, id).Error
	if err != nil {
		return nil, err
	}
	return &address, nil
}
func (r *OrderRepository) GetBookOrdersByOrderID(orderID uint) ([]entity.OrderItem, error) {
	var orderItems []entity.OrderItem
	err := r.DB.Where("order_id = ?", orderID).Preload("Book").Find(&orderItems).Error
	if err != nil {
		return nil, err
	}
	return orderItems, nil
}
func (repo OrderRepository) CreateOrder(order entity.Order) error {
	result := repo.DB.Create(&order)
	return result.Error
}

func (repo OrderRepository) UpdateOrder(id int, order entity.Order) error {
	result := repo.DB.Model(&order).Where("id = ?", id).Updates(&order)
	return result.Error
}
func GenerateSnapReq(repo *OrderRepository, orderID string, totalPrice int64) (*snap.Request, error) {
	// Get user data from database
	var user entity.User
	if err := repo.DB.Preload("Address").First(&user).Error; err != nil {
		return nil, err
	}
	var orderItems []entity.OrderItem
	if err := repo.DB.Where("order_id = ?", orderID).Find(&orderItems).Error; err != nil {
		return nil, err
	}

	// Set customer detail data from user entity
	var address entity.Address
	if len(user.Address) > 0 {
		address = user.Address[0] // take the first address
	}
	custAddress := &midtrans.CustomerAddress{
		FName:       user.Name,
		Phone:       user.Phone,
		Address:     address.Street,
		City:        address.City,
		Postcode:    address.PostalCode,
		CountryCode: address.Country,
	}
	custDetail := &midtrans.CustomerDetails{
		FName:    user.Name,
		Email:    user.Email,
		Phone:    user.Phone,
		BillAddr: custAddress,
		ShipAddr: custAddress,
	}

	// Initiate Snap Request
	var itemDetails []midtrans.ItemDetails
	for _, oi := range orderItems {
		itemDetails = append(itemDetails, midtrans.ItemDetails{
			ID:    oi.BookId,
			Price: oi.Price,
			Qty:   oi.Quantity,
			Name:  oi.Book.Title,
		})
	}
	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  orderID,
			GrossAmt: totalPrice,
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail:  custDetail,
		EnabledPayments: snap.AllSnapPaymentType,
		Items:           &itemDetails,
	}
	return snapReq, nil
}
