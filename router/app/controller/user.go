package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"member/models"
	"member/pkg/message"
	"member/router/app/content/login"
	"member/router/app/content/register"
)

func (c appController) Register(ctx *gin.Context) {
	resp := message.Response{Msg: "System error"}
	value, exist := ctx.Get("validatedBody")
	if !exist {
		resp.Msg = "Validated data missing"
		resp.ResponseBadRequest(ctx)
		return
	}

	// context用什麼格式存就得使用什麼格式取出，否則會出現panic
	req := value.(*register.Request)

	err := c.userService.Register(req)
	if err != nil {
		resp.Msg = err.Error()
		resp.ResponseBadRequest(ctx)
		return
	}

	resp.Msg = "Register successfully!"
	resp.ResponseSuccess(ctx)
	return
}

func (c appController) Login(ctx *gin.Context) {
	resp := message.Response{Msg: "System error"}
	value, exist := ctx.Get("validatedBody")
	if !exist {
		resp.Msg = "Validated data missing"
		resp.ResponseBadRequest(ctx)
		return
	}

	req := value.(*login.Request)
	jwtToken, err := c.userService.Login(req)
	if err != nil {
		resp.Msg = err.Error()
		resp.ResponseBadRequest(ctx)
		return
	}

	resp.MsgID = "LOGIN_SUCCESSFULLY"
	resp.Data = jwtToken
	resp.ResponseSuccess(ctx)
	return
}

func (c appController) GetUserInfo(ctx *gin.Context) {
	resp := message.Response{Msg: "System error"}
	userCtx, exist := ctx.Get("user")
	if !exist {
		resp.Msg = "Unauthorized!"
		resp.ResponseUnauthorized(ctx)
		log.Println("User in context is missing")
		return
	}

	user := userCtx.(*models.User)
	data, err := c.userService.GetUserInfo(user)
	if err != nil {
		resp.Msg = err.Error()
		resp.ResponseBadRequest(ctx)
		return
	}
	resp.Msg = "Get user info successfully!"
	resp.Data = data
	return
}

func (c appController) Logout(ctx *gin.Context) {
	resp := message.Response{Msg: "System error"}
	userCtx, exist := ctx.Get("user")
	if !exist {
		resp.Msg = "Unauthorized!"
		resp.ResponseUnauthorized(ctx)
		log.Println("User in context is missing")
		return
	}

	user := userCtx.(*models.User)
	err := c.userService.Logout(user)
	if err != nil {
		resp.Msg = err.Error()
		resp.ResponseBadRequest(ctx)
		return
	}
	resp.Msg = "Logout successfully!"
	resp.ResponseSuccess(ctx)
	return
}
