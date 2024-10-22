package controller

import (
	"github.com/gin-gonic/gin"
	"member/router/app/content/register"
	"member/router/repository"
	"member/router/service"
	"net/http"
)

func Register(c *gin.Context) {
	var req register.Request
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	repo := repository.NewRepository()
	api := service.NewUserService(repo)
	err = api.Register(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Register successfully!",
	})
	return
}
