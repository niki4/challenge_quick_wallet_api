package types

import "github.com/shopspring/decimal"

type Wallet struct {
	ID      int             `json:"id"`
	Balance decimal.Decimal `json:"balance"`
}

type ApiError struct {
	Error string `json:"error"`
}
