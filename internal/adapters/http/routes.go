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
	userRepo    repository.UserRepository
	userHandler handler.UserHandler
	userUsecase usecase.UserUsecase

	AuthHandler handler.AuthHandler
)

func declare() {
	userRepo = repository.UserRepository{DB: db.DbMysql}
	userUsecase = usecase.UserUsecase{Repo: userRepo}
	userHandler = handler.UserHandler{Usecase: userUsecase}

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

	seller.GET("/users", userHandler.GetAllUsers())
	seller.GET("/users/:id", userHandler.GetUser())
	seller.POST("/users", userHandler.CreateUser())
	seller.DELETE("/users/:id", userHandler.DeleteUser())
	seller.PUT("/users/:id", userHandler.UpdateUser())

	return e
}
