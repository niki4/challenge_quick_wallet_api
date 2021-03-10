package types

import (
	"github.com/shopspring/decimal"
)

type Wallet struct {
	ID      uint            `json:"id" gorm:"primaryKey"`
	Balance decimal.Decimal `json:"balance" gorm:"balance"`
}

type ApiError struct {
	Error string `json:"error"`
}
