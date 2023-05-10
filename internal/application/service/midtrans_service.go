package service

import (
	"context"
	"fmt"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/example"
	"github.com/midtrans/midtrans-go/snap"
)

var s snap.Client

func SetupGlobalMidtransConfig() {
	midtrans.ServerKey = "SB-Mid-server-x5fJwvwyo1cF2z5qGSD74Fsi"
	midtrans.Environment = midtrans.Sandbox

	// Optional : here is how if you want to set append payment notification globally
	midtrans.SetPaymentAppendNotification("https://example.com/append")
	// Optional : here is how if you want to set override payment notification globally
	midtrans.SetPaymentOverrideNotification("https://example.com/override")

	//// remove the comment bellow, in cases you need to change the default for Log Level
	// midtrans.DefaultLoggerLevel = &midtrans.LoggerImplementation{
	//	 LogLevel: midtrans.LogInfo,
	// }
}

func InitializeSnapClient() {
	s.New("SB-Mid-server-x5fJwvwyo1cF2z5qGSD74Fsi", midtrans.Sandbox)
}

func CreateTransactionWithGlobalConfig() {
	res, err := snap.CreateTransactionWithMap(example.SnapParamWithMap())
	if err != nil {
		fmt.Println("Snap Request Error", err.GetMessage())
	}
	fmt.Println("Snap response", res)
}

func CreateTransaction(snapReq snap.Request) (*snap.Response, error) {
	// Optional : here is how if you want to set append payment notification for this request
	s.Options.SetPaymentAppendNotification("https://example.com/append")

	// Optional : here is how if you want to set override payment notification for this request
	s.Options.SetPaymentOverrideNotification("https://example.com/override")
	// Send request to Midtrans Snap API

	resp, err := s.CreateTransaction(&snapReq)
	if err != nil {
		fmt.Println("Error :", err.GetMessage())
	}
	fmt.Println("Response : ", resp)
	return resp, nil
}

func CreateTokenTransactionWithGateway(snapReq snap.Request) {
	s.Options.SetPaymentOverrideNotification("https://example.com/url2")

	resp, err := s.CreateTransactionToken(&snapReq)
	if err != nil {
		fmt.Println("Error :", err.GetMessage())
	}
	fmt.Println("Response : ", resp)
}

func CreateUrlTransactionWithGateway(snapReq snap.Request) {
	s.Options.SetContext(context.Background())

	resp, err := s.CreateTransactionUrl(&snapReq)
	if err != nil {
		fmt.Println("Error :", err.GetMessage())
	}
	fmt.Println("Response : ", resp)
}
