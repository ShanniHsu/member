package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"member/router/api/content/register"
	"member/router/middleware"
	"net/http"
)

func Init() {
	r := newRouter()
	r.Run(":8888") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func newRouter() *gin.Engine {
	router := gin.New()
	// 註冊上面自定義的日誌中間件
	router.Use(middleware.Logger())
	router.GET("/test", func(c *gin.Context) {
		// 查詢我們之前在日誌中間件，注入的鍵值數
		example := c.MustGet("example").(string)
		log.Println(example)
	})

	router.GET("/ping", func(c *gin.Context) {
		panic("errrrrrr")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/register", func(c *gin.Context) {
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
	})
	return router
}
