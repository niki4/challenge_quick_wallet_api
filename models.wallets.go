package main

import (
	"errors"
	"github.com/shopspring/decimal"
)

// this is for demo only until we have not implemented DB layer
var wallets = map[int]*wallet{
	1: &wallet{1, decimal.RequireFromString("12.34")},
	2: &wallet{2, decimal.RequireFromString("56.78")},
}

// return a wallet with a specified ID
func getWalletByID(id int) (*wallet, error) {
	if w, ok := wallets[id]; ok {
		return w, nil
	}
	return nil, errors.New("wallet not found")
}

// creditWallet add a credit amount to the wallet balance
func creditWallet(id int, credit decimal.Decimal) (*wallet, error) {
	w, ok := wallets[id]
	if !ok {
		return nil, errors.New("wallet not found")
	}
	w.Balance = w.Balance.Add(credit)
	return w, nil
}

// debitWallet subtract a debit amount from the wallet balance,
// returns error in case not sufficient funds
func debitWallet(id int, debit decimal.Decimal) (*wallet, error) {
	w, ok := wallets[id]
	if !ok {
		return nil, errors.New("wallet not found")
	}
	if w.Balance.LessThan(debit) {
		return nil, errors.New("not enough money in the wallet")
	}
	w.Balance = w.Balance.Sub(debit)
	return w, nil
}
