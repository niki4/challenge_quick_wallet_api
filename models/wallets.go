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
	result := p.DB.First(w, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("wallet not found")
		}
		return nil, result.Error
	}
	return w, nil
}

// CreditWallet adds a credit amount to the wallet balance
func (p *Repository) CreditWallet(id int, credit decimal.Decimal) (*types.Wallet, error) {
	if credit.IsNegative() {
		return nil, errors.New("credit amount cannot be negative")
	}

	w := new(types.Wallet)
	result := p.DB.First(w, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("wallet not found")
		}
		return nil, result.Error
	}

	w.Balance = w.Balance.Add(credit)
	result = p.DB.Save(w)
	if result.Error != nil {
		return nil, result.Error
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
	result := p.DB.First(w, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("wallet not found")
		}
		return nil, result.Error
	}

	if w.Balance.LessThan(debit) {
		return nil, errors.New("not enough money in the wallet")
	}

	w.Balance = w.Balance.Sub(debit)
	result = p.DB.Save(w)
	if result.Error != nil {
		return nil, result.Error
	}

	return w, nil
}
