package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

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

func (handler StoreHandler) CheckTransaction() echo.HandlerFunc {
	return func(e echo.Context) error {
		orderId, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}
		trasaction, err := service.CheckTransaction(fmt.Sprint(orderId))
		return e.JSON(http.StatusOK, trasaction)
	}
}

func (handler StoreHandler) CheckStatusB2B() echo.HandlerFunc {
	return func(e echo.Context) error {
		orderId, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}
		trasaction, err := service.CheckStatusB2B(fmt.Sprint(orderId))
		return e.JSON(http.StatusOK, trasaction)
	}
}

func (handler StoreHandler) ApproveTransaction() echo.HandlerFunc {
	return func(e echo.Context) error {
		orderId, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}
		trasaction, err := service.ApproveTransaction(fmt.Sprint(orderId))
		return e.JSON(http.StatusOK, trasaction)
	}
}

func (handler StoreHandler) DenyTransaction() echo.HandlerFunc {
	return func(e echo.Context) error {
		orderId, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}
		trasaction, err := service.DenyTransaction(fmt.Sprint(orderId))
		return e.JSON(http.StatusOK, trasaction)
	}
}

func (handler StoreHandler) CancelTransaction() echo.HandlerFunc {
	return func(e echo.Context) error {
		orderId, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}
		trasaction, err := service.CancelTransaction(fmt.Sprint(orderId))
		return e.JSON(http.StatusOK, trasaction)
	}
}
func (handler StoreHandler) ExpireTransaction() echo.HandlerFunc {
	return func(e echo.Context) error {
		orderId, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}
		trasaction, err := service.ExpireTransaction(fmt.Sprint(orderId))
		return e.JSON(http.StatusOK, trasaction)
	}
}
func (handler TransactionHandler) CheckoutTransaction() echo.HandlerFunc {
	return func(e echo.Context) error {
		var order entity.Order
		var payment entity.Payment

		err := handler.OrdeUsecase.CreateOrder(order)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create user"})
		}
		// Get UserID
		user := e.Get("user").(*jwt.Token)
		claims := user.Claims.(*jwt.MapClaims)
		UserID := int((*claims)["id"].(float64))
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

		// create payment
		now := time.Now()
		currentDate := now.Format("2006-01-02")
		paymentExpiry := now.Add(30 * time.Minute).Format("2006-01-02 15:04:05")

		payment.PaymentToken = token
		payment.PaymentType = "Midtrans"
		payment.PaymentDate = currentDate
		payment.PaymentExpiry = paymentExpiry
		payment.PaymentStatus = "pending"
		payment.OrderId = orderID

		err = handler.OrdeUsecase.CreatePayment(payment)
		// Return response
		response := map[string]interface{}{
			"transaction":  snapReq,
			"tokenPayment": token,
			"redirectURL":  redirectURL,
			"totalQty":     totalQty,
		}
		return e.JSON(http.StatusOK, response)
	}
}
