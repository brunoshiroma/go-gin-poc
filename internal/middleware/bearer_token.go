package middleware

import (
	"crypto/subtle"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func BearerToken(token string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		scheme, provided, ok := strings.Cut(ctx.GetHeader("Authorization"), " ")
		if !ok || !strings.EqualFold(scheme, "Bearer") ||
			subtle.ConstantTimeCompare([]byte(provided), []byte(token)) != 1 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}

		ctx.Next()
	}
}
