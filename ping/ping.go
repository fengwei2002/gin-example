package ping

import "net/http"

import (
	"github.com/gin-gonic/gin"
)

func MainFunction() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run("localhost:8080")
}
