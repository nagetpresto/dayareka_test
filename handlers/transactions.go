package handlers

import (
	dto "BE/dto/result"
	transactionsdto "BE/dto/transactions"
	"BE/models"
	"BE/repositories"
	"strconv"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type handlerTransaction struct {
	TransactionRepository repositories.TransactionRepository
}

func HandlerTransaction(TransactionRepository repositories.TransactionRepository) *handlerTransaction {
	return &handlerTransaction{TransactionRepository}
}

func (h *handlerTransaction) PostTransaction(c echo.Context) error {
	request := new(transactionsdto.CreateTransactionRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	transaction := models.Transaction{
		UserID:    request.UserID,
		Menu:      request.Menu,
		Price:     request.Price,
		Qty:       request.Qty,
		Total:     request.Price*request.Qty,
		Payment:   request.Payment,
		CreatedAt: time.Now(),
	}

	transaction, err = h.TransactionRepository.PostTransaction(transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	transaction, _ = h.TransactionRepository.GetOneTransaction(transaction.ID)
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: transaction})
}

func (h *handlerTransaction) GetTransactionsSortedByLatest(c echo.Context) error {
	menu := c.QueryParam("menu")
	price, _ := strconv.Atoi(c.QueryParam("price"))

	transactions, err := h.TransactionRepository.GetTransactionsSortedByLatest(menu, price)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: transactions})
}