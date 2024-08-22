package router

import (
	"github.com/gin-gonic/gin"
	"member/models"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

func PostUser(r *gin.Engine) *gin.Engine {
	r.POST("/user/add", func(c *gin.Context) {
		var user models.User
		c.BindJSON(&user)
		c.JSON(200, user)
	})
	return r
}
