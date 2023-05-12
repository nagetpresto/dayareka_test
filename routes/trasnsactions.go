package routes

import (
	"BE/handlers"
	"BE/pkg/mysql"
	"BE/repositories"

	"github.com/labstack/echo/v4"
)

func TransactionRoutes(e *echo.Group) {
	TransactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(TransactionRepository)

	e.POST("/transaction", (h.PostTransaction))
	e.GET("/transaction", h.GetTransactionsSortedByLatest)
}