package main

import (
	"github.com/gin-gonic/gin"
	"gooj-backend/controller"
	"gooj-backend/repository"
	"gooj-backend/service"
	"net/http"
)

func main() {
	// 新建服务端
	backendServer := gin.Default()

	// 服务端测试路由注册
	backendServer.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ping success")
	})

	// User 模块
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	UserHandler := controller.NewUserHandler(userService)
	UserHandler.RegisterRoutes(backendServer)

	// 在 6060 端口启动
	backendServer.Run(":6060")
}
