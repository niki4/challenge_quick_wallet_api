package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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

// Response in JSON depending on request "Accept" header
// Can be enhanced to support more response formats, e.g. HTML or XML
func render(c *gin.Context, data gin.H) {
	acceptType := c.Request.Header.Get("Accept")
	contentType := c.Request.Header.Get("Content-Type")
	jsonType := "application/json"

	if (c.Request.Method == http.MethodGet && acceptType == jsonType) ||
		(c.Request.Method == http.MethodPost && contentType == jsonType) {
		c.JSON(http.StatusOK, data["payload"])
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, apiError{"incorrect request"})
	}
}
