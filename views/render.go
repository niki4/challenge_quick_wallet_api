package views

import (
	"github.com/gin-gonic/gin"
	"github.com/niki4/challenge_quick_wallet_api/types"
	"net/http"
)

// Render return response in JSON (depending on request "Accept"/"Content-Type" header).
// Can be enhanced to support more response formats, e.g. HTML or XML
func Render(c *gin.Context, data gin.H) {
	acceptType := c.Request.Header.Get("Accept")
	contentType := c.Request.Header.Get("Content-Type")
	jsonType := "application/json"

	if (c.Request.Method == http.MethodGet && acceptType == jsonType) ||
		(c.Request.Method == http.MethodPost && contentType == jsonType) {
		c.JSON(http.StatusOK, data["payload"])
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, types.ApiError{Error: "incorrect request"})
	}
}
