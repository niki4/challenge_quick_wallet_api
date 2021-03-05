package main

import (
	"errors"
	"github.com/shopspring/decimal"
)

// this is for demo only until we have not implemented DB layer
var wallets = map[int]wallet{
	1: {1, decimal.NewFromFloat(12.34)},
	2: {2, decimal.NewFromFloat(56.78)},
}

// return a wallet with a specified ID
func getWalletByID(id int) (*wallet, error) {
	if w, ok := wallets[id]; ok {
		return &w, nil
	}
	return nil, errors.New("wallet not found")
}
