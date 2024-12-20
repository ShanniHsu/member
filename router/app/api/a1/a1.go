package a1

import (
	"github.com/gin-gonic/gin"
	"member/router/app/controller"
	"member/router/app/middleware"
	"member/router/repository"
	"member/router/service"
)

func Init(router *gin.Engine) {
	newRepo := repository.NewRepository()
	newUserService := service.NewUserService(newRepo)
	api := controller.NewAppController(newUserService)
	router.POST("/register", api.Register) // 註冊
	router.POST("/login", api.Login)       // 登入

	auth := router.Group("/auth")
	auth.Use(middleware.Auth(newUserService))
	auth.GET("/info", api.GetUserInfo)          // 獲取個人資料
	auth.POST("/create-order", api.CreateOrder) // 建立訂單
}
