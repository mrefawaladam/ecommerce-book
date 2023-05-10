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
	// store management
	storeRepo    repository.StoreRepository
	storeHandler handler.StoreHandler
	storeUsecase usecase.StoreUsecase
	// auth
	AuthHandler handler.AuthHandler
	// transaction
	transactionHandler handler.TransactionHandler
	// ongkir
	ongkirHandler handler.OngkirHandler
	// order
	orderItemUsecase usecase.OrderItemsUsecase
	orderItemRepo    repository.OrderItemsRepository
	orderUsecase     usecase.OrderUsecase
	orderRepo        repository.OrderRepository
)

func declare() {
	userRepo = repository.UserRepository{DB: db.DbMysql}
	userUsecase = usecase.UserUsecase{Repo: userRepo}
	userHandler = handler.UserHandler{Usecase: userUsecase}
	// declare book management
	bookRepo = repository.BookRepository{DB: db.DbMysql}
	bookUsecase = usecase.BookUsecase{Repo: bookRepo}
	bookHandler = handler.BookHandler{BookUsecase: bookUsecase}
	// declare store management
	storeRepo = repository.StoreRepository{DB: db.DbMysql}
	storeUsecase = usecase.StoreUsecase{Repo: storeRepo}
	storeHandler = handler.StoreHandler{StoreUsecase: storeUsecase}
	// transaction
	orderRepo = repository.OrderRepository{DB: db.DbMysql}
	orderUsecase = usecase.OrderUsecase{OrderRepo: orderRepo, UserRepo: userRepo}
	orderItemRepo = repository.OrderItemsRepository{DB: db.DbMysql}
	orderItemUsecase = usecase.OrderItemsUsecase{Repo: orderItemRepo}
	transactionHandler = handler.TransactionHandler{
		Usecase:           userUsecase,
		OrderITemsUsecase: orderItemUsecase,
		OrdeUsecase:       orderUsecase}
	// Auth
	AuthHandler = handler.AuthHandler{Usecase: userUsecase}
	// order
	ongkirHandler = handler.OngkirHandler{OrderUsecase: orderUsecase}

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

	admin.GET("/stores", storeHandler.GetAllStores())
	admin.DELETE("/stores/:id", storeHandler.DeleteStore())

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

	seller.GET("/stores/my/:id", storeHandler.GetStore())
	seller.PUT("/stores/:id", storeHandler.UpdateStore())

	// seller group
	customer := e.Group("/customer")
	customer.Use(middleware.Logger())
	customer.Use(middlewares.AuthMiddleware())
	customer.Use(middlewares.RequireRole("customer"))

	// claims store
	customer.POST("/stores/claims", storeHandler.CreateStore())
	customer.GET("/books", bookHandler.GetAllBooks())
	customer.GET("/books", bookHandler.GetAllBooks())

	// transactions
	customer.POST("/trasaction/checkout", transactionHandler.CheckoutTransaction())
	customer.GET("/trasaction/check-trasaction/:id", transactionHandler.CheckTransaction())
	customer.GET("/trasaction/check-status/:id", transactionHandler.CheckStatusB2B())
	customer.GET("/trasaction/approval-trasaction/:id", transactionHandler.ApproveTransaction())
	customer.GET("/trasaction/deny-transaction/:id", transactionHandler.DenyTransaction())
	customer.GET("/trasaction/cencel-transaction/:id", transactionHandler.CancelTransaction())
	customer.GET("/trasaction/cencel-expire-transaction/:id", transactionHandler.ExpireTransaction())

	// raja onkir
	customer.POST("/ongkir/calculate", ongkirHandler.CalculateOngkir())

	// get all ebooks
	customer.GET("/books", bookHandler.GetAllBooks())

	return e
}
