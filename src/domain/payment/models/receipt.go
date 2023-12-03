package models

import "time"

type Receipt struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	OrderID   string    `json:"order_id" gorm:"index:idx_order_receipt"`
	CardID    string    `json:"card_id"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
