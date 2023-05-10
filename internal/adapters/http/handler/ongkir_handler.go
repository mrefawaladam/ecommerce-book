package handler

import (
	"fmt"
	"net/http"
	"time"

	"ebook/internal/application/usecase"

	ro "github.com/Bhinneka/go-rajaongkir"

	"github.com/labstack/echo/v4"
)

type OngkirHandler struct {
	OrderUsecase usecase.OrderUsecase
}

func (handler OngkirHandler) CalculateOngkir() echo.HandlerFunc {
	return func(e echo.Context) error {
		raja := ro.New("736cb684012893b91477f360458ea29d", 10*time.Second)
		var onkirReq struct {
			Origin      string `json:"origin"`
			Destination string `json:"shipping_address"`
			Weight      int    `json:"total_shipping"`
			Courier     string `json:"courier"`
		}
		if err := e.Bind(&onkirReq); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}
		q := ro.QueryRequest{Origin: onkirReq.Origin, Destination: onkirReq.Destination, Weight: onkirReq.Weight, Courier: onkirReq.Courier}
		result := raja.GetCost(q)

		if result.Error != nil {
			fmt.Println(result.Error.Error())
		}

		cost, ok := result.Result.(ro.Cost)
		if !ok {
			fmt.Println("Result is not Cost")
		}

		fmt.Println(cost)

		return e.JSON(http.StatusOK, cost)
	}
}
