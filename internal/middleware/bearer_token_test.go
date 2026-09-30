package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestBearerToken(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name       string
		header     string
		statusCode int
	}{
		{name: "missing token", statusCode: http.StatusUnauthorized},
		{name: "wrong scheme", header: "Basic secret-token", statusCode: http.StatusUnauthorized},
		{name: "wrong token", header: "******", statusCode: http.StatusUnauthorized},
		{name: "valid token", header: "******", statusCode: http.StatusNoContent},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			router := gin.New()
			router.GET("/", BearerToken("secret-token"), func(ctx *gin.Context) {
				ctx.Status(http.StatusNoContent)
			})

			request := httptest.NewRequest(http.MethodGet, "/", nil)
			if test.header != "" {
				request.Header.Set("Authorization", test.header)
			}
			response := httptest.NewRecorder()
			router.ServeHTTP(response, request)

			assert.Equal(t, test.statusCode, response.Code)
		})
	}
}
