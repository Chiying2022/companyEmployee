package middleware

import (
	"strings"

	"net/http"

	"git.star-link-oa.com/pkg/logger/logger"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearerToken := ctx.GetHeader("Authorization")
		if bearerToken == "" || !strings.HasPrefix(bearerToken, "Bearer ") {
			logger.FromCtx(ctx).Warnf("解析Bearer Token Header[%s] 不正確", bearerToken)
			ctx.Abort()
			return
		}

		// Remove "Bearer " prefix and validate the token
		bearerToken = strings.TrimPrefix(bearerToken, "Bearer ")
		if !isValidToken(bearerToken) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			ctx.Abort()
			return
		}

		// If valid, proceed to the next middleware/handler
		ctx.Next()

	}
}

func isValidToken(token string) bool {
	// Example token validation
	return token == "token123"
}
