package service

import (
	"fmt"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/example"
)

var c coreapi.Client

func initiateCoreApiClient() {
	c.New(example.SandboxServerKey1, midtrans.Sandbox)
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
