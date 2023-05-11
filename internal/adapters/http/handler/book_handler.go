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

type BookHandler struct {
	BookUsecase      usecase.BookUsecase
	OrderItemUsecase usecase.OrderItemsUsecase
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
		user := e.Get("user").(*jwt.Token)
		claims := user.Claims.(*jwt.MapClaims)
		sellerID := int((*claims)["id"].(float64))
		book.SellerId = uint(sellerID)

		var err error
		err = handler.BookUsecase.CreateBook(book)
		if err != nil {
			return e.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to create book"})
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
func (h BookHandler) FilterBooks() echo.HandlerFunc {
	return func(c echo.Context) error {
		// ambil parameter pencarian dari query parameter
		searchQuery := c.QueryParam("q")

		// ambil parameter filter terlaris dari query parameter
		isBestSeller, err := strconv.ParseBool(c.QueryParam("best_seller"))
		if err != nil {
			isBestSeller = false
		}

		// lakukan pencarian buku berdasarkan judul dan/atau penulis
		books, err := h.BookUsecase.SearchBooks(searchQuery)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "Internal Server Error",
			})
		}

		// filter buku yang terlaris jika diminta
		if isBestSeller {
			var orderItems []entity.OrderItem
			orderItems, err = h.OrderItemUsecase.GetBestSellingBooks()
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
					"message": "Internal Server Error",
				})
			}

			// buat map untuk menghitung jumlah buku yang terjual
			bookSoldMap := make(map[uint]int32)
			for _, orderItem := range orderItems {
				bookID, err := strconv.ParseUint(orderItem.BookId, 10, 64)
				if err != nil {
					// handle error
				}
				bookSoldMap[uint(bookID)] += orderItem.Quantity
			}

			// filter buku yang paling banyak terjual
			maxSold := int32(0)
			filteredBooks := make([]entity.Book, 0)
			for _, book := range books {
				if bookSoldMap[book.ID] > maxSold {
					maxSold = bookSoldMap[book.ID]
					filteredBooks = []entity.Book{book}
				} else if bookSoldMap[book.ID] == maxSold {
					filteredBooks = append(filteredBooks, book)
				}
			}

			books = filteredBooks
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"data": books,
		})
	}
}
