package main

import (
	"net/http"
	"sqltest/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/api/bank", listBank)
	// router.POST("/movies", createmoviesHandler)
	// router.GET("/movies/:id", getmoviesbyid)
	router.Run("localhost:8080")
}

func listBank(c *gin.Context) {
	sqltest := models.BankHandler()
	if sqltest == nil || len(sqltest) == 0 {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.IndentedJSON(http.StatusOK, sqltest)
	}
}
