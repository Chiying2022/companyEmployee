package middleware

import "github.com/gin-gonic/gin"

func CheckCondition(c *gin.Context) bool {
	// 這裡寫你的條件判斷邏輯，例如：
	return c.GetHeader("X-Special-Header") == "SpecialValue"
}
