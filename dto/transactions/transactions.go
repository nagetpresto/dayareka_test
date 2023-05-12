package transactionsdto

import (
	"time"
)

type CreateTransactionRequest struct {
	ID			int			`json:"id"`
	UserID   	int    		`json:"user_id" validate:"required"`
	Menu      	string  	`json:"menu" validate:"required"`
	Price     	int         `json:"price" validate:"required"`
	Qty       	int         `json:"qty" validate:"required"`
	Payment   	string      `json:"payment" validate:"required"`
	CreatedAt 	time.Time	`json:"created_at"`
}