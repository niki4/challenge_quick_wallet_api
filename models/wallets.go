package models

import (
	"errors"
	"github.com/niki4/challenge_quick_wallet_api/types"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

// GetWalletByID returns a wallet with a specified ID
func (p *Repository) GetWalletByID(id int) (*types.Wallet, error) {
	w := new(types.Wallet)
	p.DB.First(w, id)
	if w.ID == 0 {
		return nil, errors.New("wallet not found")
	}
	return w, nil
}

// CreditWallet adds a credit amount to the wallet balance
func (p *Repository) CreditWallet(id int, credit decimal.Decimal) (*types.Wallet, error) {
	if credit.IsNegative() {
		return nil, errors.New("credit amount cannot be negative")
	}
	w := new(types.Wallet)
	res := p.DB.Find(w, id).Update("balance", w.Balance.Add(credit))
	if res.Error != nil {
		return nil, res.Error
	}
	return w, nil
}

// DebitWallet subtracts a debit amount from the wallet balance,
// returns error in case not sufficient funds
func (p *Repository) DebitWallet(id int, debit decimal.Decimal) (*types.Wallet, error) {
	if debit.IsNegative() {
		return nil, errors.New("debit amount cannot be negative")
	}
	w := new(types.Wallet)
	p.DB.First(w, id)
	if w.ID == 0 {
		return nil, errors.New("wallet not found")
	}

	if w.Balance.LessThan(debit) {
		return nil, errors.New("not enough money in the wallet")
	}

	res := p.DB.Model(w).Update("balance", w.Balance.Sub(debit))
	if res.Error != nil {
		return nil, res.Error
	}

	return w, nil
}
