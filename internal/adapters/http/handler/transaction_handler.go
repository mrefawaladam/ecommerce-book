package handler

import (
	"net/http"

	"ebook/internal/application/service"
	"ebook/internal/application/usecase"
	"ebook/internal/entity"

	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	Usecase           usecase.UserUsecase
	OrderITemsUsecase usecase.OrderItemsUsecase
	OrdeUsecase       usecase.OrderUsecase
}

func (handler TransactionHandler) CheckoutTransaction() echo.HandlerFunc {
	return func(e echo.Context) error {
		// var order []entity.Order

		// Get UserID
		// user := e.Get("user").(*jwt.Token)
		// claims := user.Claims.(*jwt.MapClaims)
		// UserID := int((*claims)["id"].(float64))
		// var payment []entity.Payment
		orderItems := []entity.OrderItem{}
		err := e.Bind(&orderItems)
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}
		// Set order_id for each order item and save to database
		orderID, err := handler.OrdeUsecase.GetLastOrderID()

		for _, item := range orderItems {
			item.OrderId = orderID
			err = handler.OrderITemsUsecase.CreateOrderItems(item)
			if err != nil {
				return e.JSON(http.StatusInternalServerError, map[string]string{
					"error": err.Error(),
				})
			}
		}
		// Calculate totalPrice and totalQty
		totalPrice, totalQty, err := service.CalculateTotal(orderItems)
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}
		response := map[string]interface{}{
			"order_items": orderItems,
			"total_price": totalPrice,
			"total_qty":   totalQty,
		}
		return e.JSON(http.StatusOK, response)

		// // Create request body for Midtrans Snap API
		// snapReq, err := handler.OrdeUsecase.GenerateSnapReq(orderID, UserID, totalPrice)
		// if err != nil {
		// 	return nil
		// }
		// fmt.Println("================ Request with global config ================")
		// service.SetupGlobalMidtransConfig()
		// service.CreateTransactionWithGlobalConfig()

		// fmt.Println("================ Request with Snap Client ================")
		// service.InitializeSnapClient()
		// service.CreateTransaction(*snapReq)

		// fmt.Println("================ Request Snap token ================")
		// service.CreateTokenTransactionWithGateway(*snapReq)

		// fmt.Println("================ Request Snap URL ================")
		// service.CreateUrlTransactionWithGateway(*snapReq)
		// // Return response
		// response := map[string]interface{}{
		// 	"order_items": orderItems,
		// 	"total_price": totalPrice,
		// 	"total_qty":   totalQty,
		// }
		// return e.JSON(http.StatusOK, response)
	}
}
