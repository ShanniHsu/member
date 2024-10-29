package a1

import (
	"github.com/gin-gonic/gin"
	"member/router/app/controller"
	"member/router/repository"
	"member/router/service"
)

func Init(router *gin.Engine) {
	newRepo := repository.NewRepository()
	newUserService := service.NewUserService(newRepo)
	api := controller.NewAppController(newUserService)
	router.POST("/register", api.Register)
}
