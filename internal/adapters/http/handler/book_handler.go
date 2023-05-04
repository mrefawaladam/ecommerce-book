package handler

import (
	"net/http"
	"strconv"

	"ebook/internal/application/usecase"
	"ebook/internal/entity"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	BookUsecase usecase.BookUsecase
}

func (handler BookHandler) GetAllBooks() echo.HandlerFunc {
	return func(e echo.Context) error {
		var books []entity.Book

		books, err := handler.BookUsecase.GetAllBooks()
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get all books",
			"books":   books,
		})
	}
}

func (handler BookHandler) GetBook() echo.HandlerFunc {
	return func(e echo.Context) error {
		var book entity.Book
		id, err := strconv.Atoi(e.Param("id"))

		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		err = handler.BookUsecase.FindBook(id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"message": "Record Not Found",
			})
		}

		book, err = handler.BookUsecase.GetBook(id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get book",
			"book":    book,
		})
	}
}

func (handler BookHandler) CreateBook() echo.HandlerFunc {
	return func(e echo.Context) error {
		var book entity.Book
		if err := e.Bind(&book); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
		}

		validate := validator.New()
		if err := validate.Struct(book); err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{"message": "Validation errors", "errors": err.Error()})
		}
		var err error
		err = handler.BookUsecase.CreateBook(book)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create user"})
		}

		return e.JSON(http.StatusCreated, book)
	}
}

func (handler BookHandler) UpdateBook() echo.HandlerFunc {
	var book entity.Book

	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "input id is not a number",
			})
		}

		err = handler.BookUsecase.FindBook(id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"message": "Record Not Found",
			})
		}

		if err := e.Bind(&book); err != nil {
			return e.JSON(400, echo.Map{
				"error": err.Error(),
			})
		}

		err = handler.BookUsecase.UpdateBook(id, book)
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "success update book",
		})
	}
}

func (handler BookHandler) DeleteBook() echo.HandlerFunc {
	return func(e echo.Context) error {
		id, err := strconv.Atoi(e.Param("id"))
		if err != nil {
			return e.JSON(http.StatusBadRequest, map[string]interface{}{
				"messages": "input id is not a number",
			})
		}

		err = handler.BookUsecase.FindBook(id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"message": "Record Not Found",
			})
		}

		err = handler.BookUsecase.DeleteBook(id)
		if err != nil {
			return e.JSON(500, echo.Map{
				"error": err.Error(),
			})
		}

		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Delete Book`",
		})
	}
}
