package usecase

import (
	"ebook/internal/adapters/repository"
	"ebook/internal/entity"
	"fmt"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type OrderUsecase struct {
	OrderRepo repository.OrderRepository
	UserRepo  repository.UserRepository
}

func (uc *OrderUsecase) GetLastOrderID() (uint, error) {
	stores, err := uc.OrderRepo.GetLastOrderID()
	return stores, err
}

func (usecase OrderUsecase) CreateOrder(user entity.Order) error {
	err := usecase.OrderRepo.CreateOrder(user)
	return err
}

func (usecase OrderUsecase) UpdateOrder(id int, order entity.Order) error {
	err := usecase.OrderRepo.UpdateOrder(id, order)
	return err
}

// GenerateSnapReq creates a Snap Request object for generating payment token for a specific order.
func (uc *OrderUsecase) GenerateSnapReq(OrderID uint, UserID int, TotalPrice int) (*snap.Request, error) {
	// Get the order and its related data from the repository
	order, err := uc.OrderRepo.FindByID(OrderID)
	if err != nil {
		return nil, err
	}
	user, err := uc.UserRepo.GetUser(UserID)
	if err != nil {
		return nil, err
	}
	address, err := uc.UserRepo.GetAddressByID(UserID)
	if err != nil {
		return nil, err
	}
	bookOrders, err := uc.OrderRepo.GetBookOrdersByOrderID(order.ID)
	if err != nil {
		return nil, err
	}

	// Set customer detail data
	custAddress := &midtrans.CustomerAddress{
		FName:       user.Name,
		LName:       "Doe",
		Phone:       user.Phone,
		Address:     address.Street,
		City:        address.City,
		Postcode:    address.PostalCode,
		CountryCode: address.Country,
	}
	custDetail := &midtrans.CustomerDetails{
		FName:    user.Name,
		LName:    "Doe",
		Email:    user.Email,
		Phone:    user.Phone,
		BillAddr: custAddress,
		ShipAddr: custAddress,
	}

	// Create ItemDetails array for Snap Request
	var itemDetails []midtrans.ItemDetails
	var totalPrice int64 = 0
	for _, bo := range bookOrders {
		itemDetails = append(itemDetails, midtrans.ItemDetails{
			ID:    bo.BookId,
			Price: int64(bo.Book.Price),
			Qty:   int32(bo.Quantity),
			Name:  bo.Book.Title,
		})
		totalPrice += int64(bo.Book.Price) * int64(bo.Quantity)
	}

	// Check if the TotalPrice is equal to the sum of Price * Quantity
	// if totalPrice != int64(TotalPrice) {
	// 	return nil, fmt.Errorf("total price (%d) is not equal to the sum of price * quantity (%d)", totalPrice, int64(TotalPrice))
	// }

	// Create Snap Request object
	snapReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  fmt.Sprint(OrderID),
			GrossAmt: int64(totalPrice),
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
