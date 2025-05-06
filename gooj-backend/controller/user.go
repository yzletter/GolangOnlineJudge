package controller

import (
	"github.com/gin-gonic/gin"
	"gooj-backend/service"
	"net/http"
)

type UserHandler struct {
	svc *service.UserService
}

// NewUserHandler 构造函数
func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

// RegisterRoutes 注册路由
func (uh *UserHandler) RegisterRoutes(server *gin.Engine) {
	// 路由分组
	routesGroup := server.Group("/user")
	routesGroup.GET("/get/login_status", uh.GetLoginStatus) // 获取登录状态
}

// GetLoginStatus 获取当前用户登录状态
func (uh *UserHandler) GetLoginStatus(ctx *gin.Context) {
	ctx.String(http.StatusOK, "登陆成功")
	return
}
