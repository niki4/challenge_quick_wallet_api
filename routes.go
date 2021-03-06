package main

import "github.com/niki4/challenge_quick_wallet_api/handlers"

func initializeRoutes() {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/wallets/:wallet_id/balance", handlers.GetWalletBalance)
		v1.POST("/wallets/:wallet_id/credit", handlers.CreditMoneyToWallet)
		v1.POST("/wallets/:wallet_id/debit", handlers.DebitMoneyFromWallet)
	}
}
