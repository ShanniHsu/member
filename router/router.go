package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"member/router/app/api/a1"
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
	router.Use(middleware.Logger(), middleware.Cors())
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
	a1.Init(router)
	return router
}
