package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

var router *gin.Engine

func main() {
	// Set Gin to production mode. Comment the line for local debug.
	gin.SetMode(gin.ReleaseMode)

	router = gin.Default()

	// register handlers
	initializeRoutes()

	// start serving the app on 0.0.0.0:8080 (for windows "localhost:8080")
	log.Fatal(router.Run())
}
