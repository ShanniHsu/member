package controller

import (
	"github.com/gin-gonic/gin"
	"member/router/api/content/register"
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

	c.JSON(http.StatusOK, gin.H{
		"message": "Register successfully!",
	})
	return
}
