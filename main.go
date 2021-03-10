package main

import (
	"github.com/gin-gonic/gin"
	"github.com/niki4/challenge_quick_wallet_api/handlers"
	"github.com/niki4/challenge_quick_wallet_api/models"
	"log"
	"os"
)

var router *gin.Engine

func main() {
	// if DEBUG env var is not set, use production mode.
	debugMode := os.Getenv("DEBUG")
	if debugMode == "" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		log.Println("************** DEBUG MODE ENABLED **************")
	}

	router = gin.Default()

	// init storage (DB)
	conn, err := models.InitStorage()
	if err != nil {
		log.Fatalln(err)
	}

	handlers.Init(conn, router)

	// start serving the app on 0.0.0.0:8080 (for windows "localhost:8080")
	log.Fatal(router.Run())
}
