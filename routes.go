package main

import "challenge_quick_wallet_api/handlers"

func initializeRoutes() {
	router.GET("/api/v1/wallets/:wallet_id/balance", handlers.GetWalletBalance)
	router.POST("/api/v1/wallets/:wallet_id/credit", handlers.CreditMoneyToWallet)
	router.POST("/api/v1/wallets/:wallet_id/debit", handlers.DebitMoneyFromWallet)
}
