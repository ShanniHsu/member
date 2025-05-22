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
	resp := message.Response{MsgID: message.SYSTEM_ERROR}
	value, exist := ctx.Get("validatedBody")
	if !exist {
		resp.MsgID = message.VALIDATED_DATA_MISSING
		resp.ResponseBadRequest(ctx)
		return
	}

	// context用什麼格式存就得使用什麼格式取出，否則會出現panic
	req := value.(*register.Request)

	err := c.userService.Register(req)
	if err != nil {
		resp.MsgID = err.Error()
		resp.ResponseBadRequest(ctx)
		return
	}

	resp.MsgID = message.REGISTER_SUCCESSFULLY
	resp.ResponseSuccess(ctx)
	return
}

func (c appController) Login(ctx *gin.Context) {
	resp := message.Response{MsgID: message.SYSTEM_ERROR}
	value, exist := ctx.Get("validatedBody")
	if !exist {
		resp.MsgID = message.VALIDATED_DATA_MISSING
		resp.ResponseBadRequest(ctx)
		return
	}

	req := value.(*login.Request)
	jwtToken, err := c.userService.Login(req)
	if err != nil {
		resp.MsgID = err.Error()
		resp.ResponseBadRequest(ctx)
		return
	}

	resp.MsgID = message.LOGIN_SUCCESSFULLY
	resp.Data = jwtToken
	resp.ResponseSuccess(ctx)
	return
}

func (c appController) GetUserInfo(ctx *gin.Context) {
	resp := message.Response{MsgID: message.SYSTEM_ERROR}
	userCtx, exist := ctx.Get("user")
	if !exist {
		resp.MsgID = message.UNAUTHORIZED
		resp.ResponseUnauthorized(ctx)
		log.Println("User in context is missing")
		return
	}

	user := userCtx.(*models.User)
	data, err := c.userService.GetUserInfo(user)
	if err != nil {
		resp.MsgID = err.Error()
		resp.ResponseBadRequest(ctx)
		return
	}
	resp.MsgID = message.GET_USER_INFO_SUCCESSFULLY
	resp.Data = data
	return
}

func (c appController) Logout(ctx *gin.Context) {
	resp := message.Response{MsgID: message.SYSTEM_ERROR}
	userCtx, exist := ctx.Get("user")
	if !exist {
		resp.MsgID = message.UNAUTHORIZED
		resp.ResponseUnauthorized(ctx)
		log.Println("User in context is missing")
		return
	}

	user := userCtx.(*models.User)
	err := c.userService.Logout(user)
	if err != nil {
		resp.MsgID = err.Error()
		resp.ResponseBadRequest(ctx)
		return
	}
	resp.MsgID = message.LOGOUT_SUCCESSFULLY
	resp.ResponseSuccess(ctx)
	return
}
