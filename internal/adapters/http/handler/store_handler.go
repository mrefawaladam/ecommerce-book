package handler

import (
	"net/http"
	"strconv"

	"ebook/internal/application/usecase"
	"ebook/internal/entity"

	"github.com/go-playground/validator"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type StoreHandler struct {
	StoreUsecase usecase.StoreUsecase
}

func (handler StoreHandler) GetAllStores() echo.HandlerFunc {
	return func(e echo.Context) error {
		var stores []entity.Store

		stores, err := handler.StoreUsecase.GetAllStores()
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all stores",
			"stores":  stores,
		})
	}
}

func (handler StoreHandler) GetStore() echo.HandlerFunc {
	return func(e echo.Context) error {
		var store entity.Store
		id, err := strconv.Atoi(e.Param("id"))

		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		err = handler.StoreUsecase.FindStore(id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"message": "Record Not Found",
			})
		}

		store, err = handler.StoreUsecase.GetStore(id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get store",
			"store":   store,
		})
	}
}
func (handler StoreHandler) CreateStore() echo.HandlerFunc {
	return func(e echo.Context) error {
		var store entity.Store
		if err := e.Bind(&store); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
		}

		validate := validator.New()
		if err := validate.Struct(store); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Validation errors", "errors": err.Error()})
		}
		user := e.Get("user").(*jwt.Token)
		claims := user.Claims.(*jwt.MapClaims)
		sellerID := int((*claims)["id"].(float64))
		store.SellerId = uint(sellerID)

		var err error
		err = handler.StoreUsecase.CreateStore(store)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create store"})
		}

		return e.JSON(http.StatusCreated, store)
	}
}

func (handler StoreHandler) UpdateStore() echo.HandlerFunc {
	var store entity.Store

	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "input id is not a number",
			})
		}

		err = handler.StoreUsecase.FindStore(id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"message": "Record Not Found",
			})
		}

		if err := e.Bind(&store); err != nil {
			return e.JSON(400, echo.Map{
				"error": err.Error(),
			})
		}

		err = handler.StoreUsecase.UpdateStore(id, store)
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "success update store",
		})
	}
}

func (handler StoreHandler) DeleteStore() echo.HandlerFunc {
	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		err = handler.StoreUsecase.FindStore(id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"message": "Record Not Found",
			})
		}

		err = handler.StoreUsecase.DeleteStore(id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Delete Store`",
		})
	}
}
