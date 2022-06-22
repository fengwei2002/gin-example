package asciijson

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 使用 AsciiJSON 生成具有转义的非 ASCII 字符的 ASCII-only JSON。

func MainFunction() {
	r := gin.Default()

	r.GET("/asciijson", func(c *gin.Context) {
		data := map[string]any{
			"lang": "golan 封伟",
			"tag":  "<go>",
		}

		c.AsciiJSON(http.StatusOK, data)
	})

	r.Run("localhost:8080")
}
