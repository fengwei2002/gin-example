package tmplroute

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MainFunction() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/tmplindex", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "tmpl index page title",
		})

	})

	r.Run("localhost:8080")
}
