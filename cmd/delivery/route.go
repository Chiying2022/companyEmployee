package delivery

import (
	"companyEmployee/cmd/handler"
	"companyEmployee/cmd/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	engine.Use(
		cors.New(cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"GET", "POST", "PUT", "OPTIONS"},
			AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With", "X-Forwarded-For"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}),
	)

	api := engine.Group("/api")
	api.GET("/active", handler.ActiveCheck)
	api.GET("/batch_active", middleware.Authentication(), handler.ActiveBatchCheck)
}
