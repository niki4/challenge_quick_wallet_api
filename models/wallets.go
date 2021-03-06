package models

import (
	"challenge_quick_wallet_api/types"
	"errors"
	"github.com/shopspring/decimal"
)

// this is for demo only until we have not implemented DB layer
var wallets = map[int]*types.Wallet{
	1: {1, decimal.RequireFromString("12.34")},
	2: {2, decimal.RequireFromString("56.78")},
}

// GetWalletByID returns a wallet with a specified ID
func GetWalletByID(id int) (*types.Wallet, error) {
	if w, ok := wallets[id]; ok {
		return w, nil
	}
	return nil, errors.New("wallet not found")
}

// CreditWallet adds a credit amount to the wallet balance
func CreditWallet(id int, credit decimal.Decimal) (*types.Wallet, error) {
	if credit.IsNegative() {
		return nil, errors.New("credit amount cannot be negative")
	}
	w, ok := wallets[id]
	if !ok {
		return nil, errors.New("wallet not found")
	}
	w.Balance = w.Balance.Add(credit)
	return w, nil
}

// DebitWallet subtracts a debit amount from the wallet balance,
// returns error in case not sufficient funds
func DebitWallet(id int, debit decimal.Decimal) (*types.Wallet, error) {
	if debit.IsNegative() {
		return nil, errors.New("debit amount cannot be negative")
	}
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
