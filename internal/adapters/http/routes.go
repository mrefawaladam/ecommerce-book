package http

import (
	db "ebook/internal/adapters/db/mysql"
	handler "ebook/internal/adapters/http/handler"
	middlewares "ebook/internal/adapters/http/middleware"
	repository "ebook/internal/adapters/repository"
	usecase "ebook/internal/application/usecase"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	// user management
	userRepo    repository.UserRepository
	userHandler handler.UserHandler
	userUsecase usecase.UserUsecase
	// book management
	bookRepo    repository.BookRepository
	bookHandler handler.BookHandler
	bookUsecase usecase.BookUsecase
	// auth
	AuthHandler handler.AuthHandler
)

func declare() {
	userRepo = repository.UserRepository{DB: db.DbMysql}
	userUsecase = usecase.UserUsecase{Repo: userRepo}
	userHandler = handler.UserHandler{Usecase: userUsecase}
	// declare book management
	bookRepo = repository.BookRepository{DB: db.DbMysql}
	bookUsecase = usecase.BookUsecase{Repo: bookRepo}
	bookHandler = handler.BookHandler{BookUsecase: bookUsecase}

	AuthHandler = handler.AuthHandler{Usecase: userUsecase}
}

func InitRoutes() *echo.Echo {
	db.Init()
	declare()

	e := echo.New()
	e.POST("/login", AuthHandler.Login())
	e.POST("/registrasi", AuthHandler.Register())

	// admin group
	admin := e.Group("/admin")
	admin.Use(middleware.Logger())
	admin.Use(middlewares.AuthMiddleware())
	admin.Use(middlewares.RequireRole("admin"))

	admin.GET("/users", userHandler.GetAllUsers())
	admin.GET("/users/:id", userHandler.GetUser())
	admin.POST("/users", userHandler.CreateUser())
	admin.DELETE("/users/:id", userHandler.DeleteUser())
	admin.PUT("/users/:id", userHandler.UpdateUser())

	// seller group
	seller := e.Group("/seller")
	seller.Use(middleware.Logger())
	seller.Use(middlewares.AuthMiddleware())
	seller.Use(middlewares.RequireRole("seller"))

	seller.GET("/books", bookHandler.GetAllBooks())
	seller.GET("/books/:id", bookHandler.GetBook())
	seller.POST("/books", bookHandler.CreateBook())
	seller.DELETE("/books/:id", bookHandler.DeleteBook())
	seller.PUT("/books/:id", bookHandler.UpdateBook())

	return e
}
