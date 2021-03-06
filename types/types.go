package types

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	ID      uint            `json:"id" gorm:"primaryKey"`
	Balance decimal.Decimal `json:"balance" gorm:"balance"`
}

type ApiError struct {
	Error string `json:"error"`
}
