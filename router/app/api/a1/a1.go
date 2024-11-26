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
	auth.GET("/info/:id", api.GetUserInfo) // 獲取個人資料
	//auth.GET("/info", func(ctx *gin.Context) {
	//	claim, err := jwt.GetUserInfo(ctx)
	//	if err != nil {
	//		ctx.JSON(http.StatusBadRequest, gin.H{
	//			"message": err,
	//		})
	//		return
	//	}
	//
	//	ctx.JSON(http.StatusOK, gin.H{
	//		"message": claim,
	//	})
	//	return
	//})
}
