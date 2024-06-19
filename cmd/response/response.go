package response

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int            `json:"code"`
	Data interface{}    `json:"data"`
	M    map[string]int `json:"m"`
}

func SuccResponse(c *gin.Context, statusCode int, data interface{}, m map[string]int) {
	c.JSON(statusCode, Response{
		Code: statusCode,
		Data: data,
		M:    m,
	})
}

func FailResponse(c *gin.Context, statusCode int, err error, msg string) {
	err = errors.New(msg)
	c.AbortWithStatusJSON(statusCode, gin.H{"error": err.Error()})
}
