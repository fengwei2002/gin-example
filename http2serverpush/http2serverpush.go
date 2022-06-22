package http2serverpush

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
什么是 http2 server push

HTTP/2 服务器推送由 PUSH_PROMISE 帧驱动。
PUSH_PROMISE 描述了服务器预测浏览器将在不久的将来发出的请求。
一旦浏览器收到一个 PUSH_PROMISE，它就知道服务器将交付资源。
如果浏览器稍后发现它需要此资源，它将等待推送完成而不是发送新请求。
这减少了浏览器在网络上等待的时间。
*/

var html = template.Must(template.New("https").Parse(`
<html>
	<head>
		<title>Https Test</title>
		<script src="/assets/app.js"></script>
	</head>

	<body>
		<h1 style="color:red;">Welcome, Ginner!</h1>
	</body>
</html>
`))

func MainFunction() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.SetHTMLTemplate(html)

	r.GET("/", func(c *gin.Context) {
		if pusher := c.Writer.Pusher(); pusher != nil {
			// 先获取 pusher
			// 然后使用 pusher.Push() 做服务器推送
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}
		c.HTML(http.StatusOK, "https", gin.H{
			"status": "success",
		})
	})

	// 监听并在 https://127.0.0.1:8080 上启动服务
	/*
	TLS: SSL 的进化版本，现在被广泛使用的 https 传输层安全性协议
	*/
	r.RunTLS(":8080", "./testdata/server.pem", "./testdata/server.key")
}
