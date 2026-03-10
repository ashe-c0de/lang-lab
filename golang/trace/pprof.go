package main

import (
	"net/http"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
)

var memoryLeak [][]byte

func main() {

	// 启动 pprof server
	go func() {
		// 仅允许本机访问
		http.ListenAndServe("127.0.0.1:6060", nil)
	}()

	r := gin.Default()

	// 正常接口
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 模拟内存增长
	r.GET("/leak", func(c *gin.Context) {

		data := make([]byte, 10*1024*1024) // 10MB

		memoryLeak = append(memoryLeak, data)

		c.JSON(200, gin.H{
			"msg":   "memory allocated",
			"total": len(memoryLeak),
		})
	})

	r.Run(":8080")
}
