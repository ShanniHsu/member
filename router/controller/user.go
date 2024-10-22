package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"member/models"
	"member/router/api/content/register"
	"member/router/repository"
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

	api := repository.NewRepository().UserRepository
	resp, err := api.GetUserByAccount(models.User{}, req.Account)
	fmt.Println("resp: ", resp, ";", "err: ", err)

	c.JSON(http.StatusOK, gin.H{
		"message": "Register successfully!",
	})
	return
}
