package jsonp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MainFunction() {
	r := gin.Default()

	r.GET("/JSONP", func(c *gin.Context) {
		data := map[string]any{
			"foo": "bar",
		}

		c.JSONP(http.StatusOK, data)
	})

	r.Run(":8080")
}

/*

使用 json 可以将所有的 data 中的内容转换为 json 返回

json 和 jsonp 的区别

所谓的jsonp，就是一句函数调用，数据都被包裹传递到参数中了

为了跨域传输 json 数据

json with padding 就是把 json 对象用符合 js 语法的形式，包裹起来使得其他网站也可以访问的到
就是将 json 文件封装成 js 文件

jquery 已经将 jsonp 的功能封装在 ajax 中
*/
