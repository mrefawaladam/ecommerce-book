package http

import (
	db "ebook/internal/adapters/db/mysql"
	handler "ebook/internal/adapters/http/handler"
	repository "ebook/internal/adapters/repository"
	usecase "ebook/internal/application/usecase"

	"github.com/labstack/echo/v4"
)

var (
	userRepo    repository.Repository
	userHandler handler.Handler
	userUsecase usecase.Usecase
)

func declare() {
	userRepo = repository.Repository{DB: db.DbMysql}
	userUsecase = usecase.Usecase{Repo: userRepo}
	userHandler = handler.Handler{Usecase: userUsecase}

}

func InitRoutes() *echo.Echo {
	db.Init()
	declare()

	e := echo.New()
	user := e.Group("/users")
	user.GET("", userHandler.GetAllUsers())
	user.GET("/:id", userHandler.GetUser())
	user.POST("", userHandler.CreateUser())
	user.DELETE("/:id", userHandler.DeleteUser())
	user.PUT("/:id", userHandler.UpdateUser())

	return e
}
