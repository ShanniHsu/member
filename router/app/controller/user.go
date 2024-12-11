package controller

import (
	"github.com/gin-gonic/gin"
	"member/router/app/content/login"
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

func (c appController) Login(ctx *gin.Context) {
	req := new(login.Request)
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if req.Account == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "The account must input.",
		})
		return
	}

	if req.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "The password must input.",
		})
		return
	}
	jwtToken, err := c.userService.Login(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Login successfully!",
		"jwtToken": jwtToken,
	})
	return
}

func (c appController) GetUserInfo(ctx *gin.Context) {
	resp, err := c.userService.GetUserInfo(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get Successfully",
		"data":    resp,
	})
	return
}
