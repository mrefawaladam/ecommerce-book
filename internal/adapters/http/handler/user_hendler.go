package handler

import (
	"net/http"
	"strconv"

	"ebook/internal/application/usecase"
	"ebook/internal/entity"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	Usecase usecase.Usecase
}

func (handler Handler) GetAllUsers() echo.HandlerFunc {
	return func(e echo.Context) error {
		var users []entity.User

		users, err := handler.Usecase.GetAllUsers()
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all users",
			"users":   users,
		})
	}
}

func (handler Handler) GetUser() echo.HandlerFunc {
	return func(e echo.Context) error {
		var user entity.User
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		err = handler.Usecase.SearchUser(id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"message": "Record Not Found",
			})
		}

		user, err = handler.Usecase.GetUser(id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get user",
			"user":    user,
		})
	}
}

func (handler Handler) CreateUser() echo.HandlerFunc {
	var user entity.User

	return func(e echo.Context) error {
		if err := e.Bind(&user); err != nil {
			return e.JSON(400, echo.Map{
				"error": err.Error(),
			})
		}

		var err error

		err = handler.Usecase.CreateUser(user)
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "success create new user",
			"user":    &user,
		})
	}
}

func (handler Handler) UpdateUser() echo.HandlerFunc {
	var user entity.User

	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "input id is not a number",
			})
		}

		err = handler.Usecase.SearchUser(id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"message": "Record Not Found",
			})
		}

		if err := e.Bind(&user); err != nil {
			return e.JSON(400, echo.Map{
				"error": err.Error(),
			})
		}

		err = handler.Usecase.UpdateUser(id, user)
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "success update user",
		})
	}
}

func (handler Handler) DeleteUser() echo.HandlerFunc {
	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		err = handler.Usecase.SearchUser(id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"message": "Record Not Found",
			})
		}

		err = handler.Usecase.DeleteUser(id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Delete User`",
		})
	}
}
