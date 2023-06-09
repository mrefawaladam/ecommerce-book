package service

import (
	"fmt"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

var c coreapi.Client

func initiateCoreApiClient() {
	c.New("SB-Mid-server-x5fJwvwyo1cF2z5qGSD74Fsi", midtrans.Sandbox)
}

func CheckTransaction(OrderID string) (*coreapi.TransactionStatusResponse, error) {
	res, err := c.CheckTransaction(OrderID)
	if err != nil {
		// do something on error handle
	}
	fmt.Println("Response: ", res)
	return res, err
}

func CheckStatusB2B(OrderID string) (*coreapi.TransactionStatusB2bResponse, error) {
	res, err := c.GetStatusB2B(OrderID)
	if err != nil {
		// do something on error handle
	}
	fmt.Println("Response: ", res)
	return res, err
}

func ApproveTransaction(OrderID string) (*coreapi.ChargeResponse, error) {
	res, err := c.ApproveTransaction(OrderID)
	if err != nil {
		// do something on error handle
		return nil, err
	}
	fmt.Println("Response: ", res)
	return res, err
}

func DenyTransaction(OrderID string) (*coreapi.ChargeResponse, error) {
	res, err := c.DenyTransaction(OrderID)
	if err != nil {
		// do something on error handle
	}
	fmt.Println("Response: ", res)
	return res, err
}

func CancelTransaction(OrderID string) (*coreapi.ChargeResponse, error) {
	res, err := c.CancelTransaction(OrderID)
	if err != nil {
		// do something on error handle
	}
	fmt.Println("Response: ", res)
	return res, err
}

func ExpireTransaction(OrderID string) (*coreapi.ChargeResponse, error) {
	res, err := c.ExpireTransaction(OrderID)
	if err != nil {
		// do something on error handle
	}
	fmt.Println("Response: ", res)
	return res, err
}
