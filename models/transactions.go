package models

import (
	"time"
)

type Transaction struct {
	ID        int                   `json:"id"`
	UserID    int                   `json:"user_id"`
	User      User   				`json:"user"`
	Menu      string                `json:"menu" gorm:"type:varchar(255)"`
	Price     int                   `json:"price" gorm:"type:int"`
	Qty       int                   `json:"qty" gorm:"type:int"`
	Total     int                   `json:"total" gorm:"type:int"`
    Payment   string                `json:"payment" gorm:"type:varchar(255)"`
	CreatedAt time.Time             `json:"created_at" gorm:"type:timestamp"`
}