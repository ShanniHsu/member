package a1

import (
	"github.com/gin-gonic/gin"
	"member/pkg/jwt"
	"member/router/app/controller"
	"member/router/app/middleware"
	"member/router/repository"
	"member/router/service"
	"net/http"
)

func Init(router *gin.Engine) {
	newRepo := repository.NewRepository()
	newUserService := service.NewUserService(newRepo)
	api := controller.NewAppController(newUserService)
	router.POST("/register", api.Register)
	router.POST("/login", api.Login)

	auth := router.Group("/auth")
	auth.Use(middleware.Auth())
	auth.GET("/info", func(ctx *gin.Context) {
		claim, err := jwt.GetUserInfo(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": claim,
		})
		return
	})
}
