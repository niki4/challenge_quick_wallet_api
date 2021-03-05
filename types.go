package main

import "github.com/shopspring/decimal"

type wallet struct {
	ID      int             `json:"id"`
	Balance decimal.Decimal `json:"balance"`
}

type apiError struct {
	Error string `json:"error"`
}
