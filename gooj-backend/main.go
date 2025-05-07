package main

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/yzletter/XtremeGin/middlewares/ratelimit"
	limiter "github.com/yzletter/XtremeGin/middlewares/ratelimit/limiter/slide_window_limiter"
	"gooj-backend/controller"
	"gooj-backend/repository"
	"gooj-backend/service"
	"net/http"
	"time"
)

func main() {
	// 新建服务端
	backendServer := gin.Default()

	// Redis 服务
	rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})

	// GoOJ 限流服务
	goOJLimiter := limiter.NewRedisSlideWindowLimiter(rdb, time.Minute, 10)
	goOJRateLimitHandler := ratelimit.NewRateLimitBuilder(goOJLimiter).Build()
	backendServer.Use(goOJRateLimitHandler)

	// User 模块
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	UserHandler := controller.NewUserHandler(userService)
	UserHandler.RegisterRoutes(backendServer)

	// 服务端测试路由注册
	backendServer.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "ping success")
	})

	// 在 6060 端口启动
	backendServer.Run(":6060")
}
