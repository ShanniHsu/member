package a1

import (
	"github.com/gin-gonic/gin"
	"member/router/app/controller"
)

func Init(router *gin.Engine) {
	router.POST("/register", controller.Register)
}
