package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 新建服务端
	backendServer := gin.Default()

	// 服务端路由注册
	backendServer.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ping success")
	})

	// 在 6060 端口启动
	backendServer.Run(":6060")
}
