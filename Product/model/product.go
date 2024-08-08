package model

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	Price         uint      `json:"price"`
	StockQuantity uint      `json:"stock_quantity"`
	Description   string    `json:"discription"`
	AddedOn       time.Time `json:"added_on"`
}
