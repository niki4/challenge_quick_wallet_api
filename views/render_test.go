package views

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"net/http"
	"net/http/httptest"
	"testing"
)

var testPayload = map[string]interface{}{
	"payload": map[string]interface{}{
		"id":      123,
		"balance": decimal.RequireFromString("7.89"),
	},
}

func performRequest(r http.Handler, method, path string, headers map[string]string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func handleRenderFunc(c *gin.Context) {
	Render(c, testPayload)
}

func TestRenderJSON(t *testing.T) {
	router := gin.Default()
	// register test route with the "Render" function to test
	router.GET("/", handleRenderFunc)
	router.POST("/", handleRenderFunc)

	testCases := []struct {
		name, method string
		headers      map[string]string
		expStatus    int
	}{
		{"JSONValidContentType", "GET", map[string]string{"Accept": "application/json"}, http.StatusOK},
		{"JSONMissedContentType", "GET", map[string]string{}, http.StatusBadRequest},
		{"JSONValidContentType", "POST", map[string]string{"Content-Type": "application/json"}, http.StatusOK},
		{"JSONMissedContentType", "POST", map[string]string{}, http.StatusBadRequest},
		{"Not supported method", "PUT", map[string]string{}, http.StatusNotFound},
	}

	// Perform requests to the above defined handler
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s: %s", tc.method, tc.name), func(t *testing.T) {
			w := performRequest(router, tc.method, "/", tc.headers)
			assert.Equal(t, tc.expStatus, w.Code)
		})
	}
}
