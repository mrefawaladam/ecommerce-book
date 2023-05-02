package handler

import (
	"net/http"

	"ebook/internal/adapters/http/middleware"
	"ebook/internal/application/service"
	"ebook/internal/application/usecase"

	"ebook/internal/entity"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	Usecase usecase.Usecase
}

func (handler AuthHandler) Register() echo.HandlerFunc {
	return func(e echo.Context) error {
		var user entity.User
		if err := e.Bind(&user); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
		}

		// Validasi input menggunakan package validator
		validate := validator.New()
		if err := validate.Struct(user); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Validation errors", "errors": err.Error()})
		}

		// Validasi email unik
		if err := handler.Usecase.UniqueEmail(user.Email); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
		}

		hashedPassword, err := service.Encrypt(user.Password)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create user"})
		}
		user.Password = string(hashedPassword)
		user.Role = "customer"

		err = handler.Usecase.CreateUser(user)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create user"})
		}

		return e.JSON(http.StatusCreated, user)
	}
}

func (handler AuthHandler) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Bind request body to user struct
		var user entity.User
		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
		}

		// Get user by email
		dbUser, err := handler.Usecase.GetUserByEmail(user.Email)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid email or password"})
		}

		// Check password
		if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid email or password"})
		}

		t, err := middleware.CreateToken(int(dbUser.ID), dbUser.Email, dbUser.Role)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create token"})
		}

		return c.JSON(http.StatusOK, map[string]string{"token": t})
	}
}
