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

}

func debitMoneyFromWallet(c *gin.Context) {

}
