package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/niki4/challenge_quick_wallet_api/types"
	"github.com/niki4/challenge_quick_wallet_api/views"
	"net/http"
	"strconv"
)

// GetWalletBalance handles requests for Wallet with given ID.
// If wallet found - passes current balance to render function, if not found - abort request with status.
func (e *Env) GetWalletBalance(c *gin.Context) {
	// Check if the Wallet ID is valid
	if walletID, err := strconv.Atoi(c.Param("wallet_id")); err == nil {
		// Check if the Wallet exist
		if wallet, err := e.Repository.GetWalletByID(walletID); err == nil {
			views.Render(c, gin.H{"payload": map[string]interface{}{
				"id":      wallet.ID,
				"balance": wallet.Balance,
			}})
		} else {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
}

// CreditMoneyToWallet handles requests to credit Wallet with given ID for specified amount.
// If wallet found - passes updated balance to render function, if not found - abort request with status.
func (e *Env) CreditMoneyToWallet(c *gin.Context) {
	creditW := new(types.Wallet)
	if err := c.BindJSON(creditW); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	walletID, err := strconv.Atoi(c.Param("wallet_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if wallet, err := e.Repository.CreditWallet(walletID, creditW.Balance); err == nil {
		views.Render(c, gin.H{"payload": map[string]interface{}{
			"id":      wallet.ID,
			"balance": wallet.Balance,
		}}) // return updated wallet on success case
	} else {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError, types.ApiError{Error: err.Error()}) // e.g., credit balance failed
		return
	}
}

// DebitMoneyFromWallet handles requests to debit Wallet with given ID for specified amount.
// If wallet found - passes updated balance to render function, if not found - abort request with status.
func (e *Env) DebitMoneyFromWallet(c *gin.Context) {
	debitW := new(types.Wallet)
	if err := c.BindJSON(debitW); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	walletID, err := strconv.Atoi(c.Param("wallet_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if wallet, err := e.Repository.DebitWallet(walletID, debitW.Balance); err == nil {
		views.Render(c, gin.H{"payload": map[string]interface{}{
			"id":      wallet.ID,
			"balance": wallet.Balance,
		}}) // return updated wallet on success case
	} else {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError, types.ApiError{Error: err.Error()}) // e.g., debit balance failed
		return
	}
}
