package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"member/models"
	"member/router/app/content/login"
	"member/router/app/content/register"
	"net/http"
)

func (c appController) Register(ctx *gin.Context) {
	value, exist := ctx.Get("validatedBody")
	if !exist {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Validated data missing",
		})
		return
	}

	// context用什麼格式存就得使用什麼格式取出，否則會出現panic
	req := value.(*register.Request)

	err := c.userService.Register(req)
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
	value, exist := ctx.Get("validatedBody")
	if !exist {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Validated data missing",
		})
		return
	}

	req := value.(*login.Request)
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
	
	userCtx, exist := ctx.Get("user")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized!",
		})
		log.Print("User in context is missing")
		return
	}

	user := userCtx.(*models.User)
	resp, err := c.userService.GetUserInfo(user)
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
	userCtx, exist := ctx.Get("user")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized!",
		})
		log.Println("User in context is missing")
		return
	}

	user := userCtx.(*models.User)
	err := c.userService.Logout(user)
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
