package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/niki4/challenge_quick_wallet_api/models"
	"github.com/niki4/challenge_quick_wallet_api/types"
	"github.com/niki4/challenge_quick_wallet_api/views"
	"net/http"
	"strconv"
)

func GetWalletBalance(c *gin.Context) {
	// Check if the Wallet ID is valid
	if walletID, err := strconv.Atoi(c.Param("wallet_id")); err == nil {
		// Check if the Wallet exist
		if wallet, err := models.GetWalletByID(walletID); err == nil {
			views.Render(c, gin.H{"payload": wallet})
		} else {
			c.AbortWithStatus(http.StatusNotFound)
		}
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}

func CreditMoneyToWallet(c *gin.Context) {
	creditW := new(types.Wallet)
	if err := c.BindJSON(creditW); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	walletID, err := strconv.Atoi(c.Param("wallet_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	if wallet, err := models.CreditWallet(walletID, creditW.Balance); err == nil {
		views.Render(c, gin.H{"payload": wallet}) // return updated wallet on success case
	} else {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError, types.ApiError{Error: err.Error()}) // e.g., credit balance failed
	}
}

func DebitMoneyFromWallet(c *gin.Context) {
	debitW := new(types.Wallet)
	if err := c.BindJSON(debitW); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	walletID, err := strconv.Atoi(c.Param("wallet_id"))
	if err == nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	if wallet, err := models.DebitWallet(walletID, debitW.Balance); err == nil {
		views.Render(c, gin.H{"payload": wallet}) // return updated wallet on success case
	} else {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError, types.ApiError{Error: err.Error()}) // e.g., debit balance failed
	}
}
