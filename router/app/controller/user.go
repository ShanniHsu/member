package controller

import (
	"github.com/gin-gonic/gin"
	"member/router/app/content/login"
	"member/router/app/content/register"
	"member/router/app/content/response"
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
	res := response.Response{Message: "Server error"}
	if err != nil {
		res.Message = err.Error()
		res.ResponseBadRequest(ctx)
		return
	}
	if req.Account == "" {
		res.Message = "The account must input."
		res.ResponseBadRequest(ctx)
		return
	}

	if req.Password == "" {
		res.Message = "The password must input."
		res.ResponseBadRequest(ctx)
		return
	}
	jwtToken, err := c.userService.Login(req)
	if err != nil {
		res.Message = err.Error()
		res.ResponseBadRequest(ctx)
		return
	}
	res.Message = "Login successfully"
	res.Result = true
	res.Data = jwtToken
	res.ResponseSuccess(ctx)
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

func (c appController) Logout(ctx *gin.Context) {
	err := c.userService.Logout(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Logout Successfully",
	})
	return
}
