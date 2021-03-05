package main

func initializeRoutes() {
	router.GET("/api/v1/wallets/:wallet_id/balance", getWalletBalance)
	router.POST("/api/v1/wallets/:wallet_id/credit", creditMoneyToWallet)
	router.POST("/api/v1/wallets/:wallet_id/debit", debitMoneyFromWallet)
}
