package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/niki4/challenge_quick_wallet_api/models"
	"gorm.io/gorm"
)

type Env struct {
	Repository models.Repository
	router     *gin.Engine
}

func Init(conn *gorm.DB, router *gin.Engine) {
	env := &Env{
		Repository: models.CreateRepository(conn),
		router:     router,
	}

	// register handlers
	initializeRoutes(env)
}

func initializeRoutes(e *Env) {
	v1 := e.router.Group("/api/v1")
	{
		v1.GET("/wallets/:wallet_id/balance", e.GetWalletBalance)
		v1.POST("/wallets/:wallet_id/credit", e.CreditMoneyToWallet)
		v1.POST("/wallets/:wallet_id/debit", e.DebitMoneyFromWallet)
	}
}
