package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getWalletBalance(c *gin.Context) {
	// Check if the Wallet ID is valid
	if walletID, err := strconv.Atoi(c.Param("wallet_id")); err == nil {
		// Check if the Wallet exist
		if wallet, err := getWalletByID(walletID); err == nil {
			render(c, gin.H{"payload": wallet})
		} else {
			c.AbortWithStatus(http.StatusNotFound)
		}
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}

func creditMoneyToWallet(c *gin.Context) {
	creditW := new(wallet)
	if err := c.BindJSON(creditW); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	walletID, err := strconv.Atoi(c.Param("wallet_id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	if wallet, err := creditWallet(walletID, creditW.Balance); err == nil {
		render(c, gin.H{"payload": wallet}) // return updated wallet on success case
	} else {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError, apiError{err.Error()}) // e.g., credit balance failed
	}
}

func debitMoneyFromWallet(c *gin.Context) {
	debitW := new(wallet)
	if err := c.BindJSON(debitW); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	walletID, err := strconv.Atoi(c.Param("wallet_id"))
	if err == nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	if wallet, err := debitWallet(walletID, debitW.Balance); err == nil {
		render(c, gin.H{"payload": wallet}) // return updated wallet on success case
	} else {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError, apiError{err.Error()}) // e.g., debit balance failed
	}
}
