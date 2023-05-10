package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"ebook/internal/application/service"
	"ebook/internal/application/usecase"
	"ebook/internal/entity"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	Usecase           usecase.UserUsecase
	OrderITemsUsecase usecase.OrderItemsUsecase
	OrdeUsecase       usecase.OrderUsecase
}

func (handler TransactionHandler) CheckoutTransactionss() echo.HandlerFunc {
	return func(e echo.Context) error {

		var orderItemsReq struct {
			OrderItems []entity.OrderItem `json:"order_items"`
		}
		if err := e.Bind(&orderItemsReq); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}
		return e.JSON(http.StatusOK, orderItemsReq.OrderItems)
	}
}

func (handler TransactionHandler) CheckoutTransaction() echo.HandlerFunc {
	return func(e echo.Context) error {
		var order entity.Order
		err := handler.OrdeUsecase.CreateOrder(order)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create user"})
		}
		// Get UserID
		user := e.Get("user").(*jwt.Token)
		claims := user.Claims.(*jwt.MapClaims)
		UserID := int((*claims)["id"].(float64))
		// var payment []entity.Payment
		var orderItemsReq struct {
			OrderItems      []entity.OrderItem `json:"order_items"`
			ShippingAddress string             `json:"shipping_address"`
			TotalShipping   string             `json:"total_shipping"`
		}
		if err := e.Bind(&orderItemsReq); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}

		// Set order_id for each order item and save to database
		orderID, err := handler.OrdeUsecase.GetLastOrderID()

		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}
		for _, item := range orderItemsReq.OrderItems {
			item.OrderId = orderID
			err = handler.OrderITemsUsecase.CreateOrderItems(item)
			if err != nil {
				return e.JSON(http.StatusInternalServerError, map[string]string{
					"error": err.Error(),
				})
			}
		}
		// Calculate totalPrice and totalQty
		totalPrice, totalQty, err := service.CalculateTotal(orderItemsReq.OrderItems)
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}
		// declare variables order
		order.ShippingAddress = orderItemsReq.ShippingAddress
		order.TotalShipping = orderItemsReq.TotalShipping
		order.CustomerID = uint(UserID)
		order.TotalOrder = strconv.Itoa(totalPrice)

		err = handler.OrdeUsecase.UpdateOrder(int(orderID), order)
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		// Create request body for Midtrans Snap API
		snapReq, err := handler.OrdeUsecase.GenerateSnapReq(orderID, UserID, totalPrice)
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}
		fmt.Println("================ Request with global config ================")
		service.SetupGlobalMidtransConfig()
		service.CreateTransactionWithGlobalConfig()

		fmt.Println("================ Request with Snap Client ================")
		service.InitializeSnapClient()
		respPayment, err := service.CreateTransaction(*snapReq)
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}
		token := respPayment.Token
		redirectURL := respPayment.RedirectURL

		// Return response
		response := map[string]interface{}{
			"transaction":  snapReq,
			"tokenPayment": token,
			"redirectURL":  redirectURL,
		}
		return e.JSON(http.StatusOK, response)
	}
}
