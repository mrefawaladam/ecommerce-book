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
	userRepo    repository.Repository
	userHandler handler.HandlerUser
	userUsecase usecase.Usecase

	AuthHandler handler.AuthHandler
)

func declare() {
	userRepo = repository.Repository{DB: db.DbMysql}
	userUsecase = usecase.Usecase{Repo: userRepo}
	userHandler = handler.HandlerUser{Usecase: userUsecase}
	AuthHandler = handler.AuthHandler{Usecase: userUsecase}
}

func InitRoutes() *echo.Echo {
	db.Init()
	declare()

	e := echo.New()
	e.POST("/login", AuthHandler.Login())
	e.POST("/registrasi", AuthHandler.Register())

	admin := e.Group("/admin")
	admin.Use(middleware.Logger())
	admin.Use(middlewares.AuthMiddleware())
	admin.Use(middlewares.RequireRole("admin"))

	admin.GET("/users", userHandler.GetAllUsers())
	admin.GET("/users/:id", userHandler.GetUser())
	admin.POST("/users", userHandler.CreateUser())
	admin.DELETE("/users/:id", userHandler.DeleteUser())
	admin.PUT("/users/:id", userHandler.UpdateUser())

	return e
}
