package main

import (
	"errors"
	"github.com/shopspring/decimal"
)

// this is for demo only until we have not implemented DB layer
var walletList = []wallet{
	{1, decimal.NewFromFloat(12.34)},
	{2, decimal.NewFromFloat(56.78)},
}

// return a wallet with a specified ID
func getWalletByID(id int) (*wallet, error) {
	for _, w := range walletList {
		if w.ID == id {
			return &w, nil
		}
	}
	return nil, errors.New("wallet not found")
}
