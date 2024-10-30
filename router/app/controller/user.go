package controller

import (
	"github.com/gin-gonic/gin"
	"member/router/app/content/register"
	"net/http"
)

func (c appController) Register(ctx *gin.Context) {
	req := new(register.Request)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = c.userService.Register(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Register successfully!",
	})
	return
}
