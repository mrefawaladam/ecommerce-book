package http

import (
	db "ebook/internal/adapters/db/mysql"
	handler "ebook/internal/adapters/http/handler"
	middlewares "ebook/internal/adapters/http/middleware"
	repository "ebook/internal/adapters/repository"
	usecase "ebook/internal/application/usecase"

	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
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

	// middleware
	e.Use(middleware.Logger())
	// protected route

	protected := e.Group("/protected")
	protected.Use(middlewares.AuthMiddleware())
	protected.GET("", func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*jwt.MapClaims)
		userID := int((*claims)["id"].(float64))
		return c.String(http.StatusOK, fmt.Sprintf("User ID: %d", userID))
	})

	user := e.Group("/users")
	user.GET("", userHandler.GetAllUsers())
	user.GET("/:id", userHandler.GetUser())
	user.POST("", userHandler.CreateUser())
	user.DELETE("/:id", userHandler.DeleteUser())
	user.PUT("/:id", userHandler.UpdateUser())

	return e
}
