package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"member/router/middleware"
	"net/http"
)

func Init() {
	//r := gin.Default()
	r := gin.New()
	// 註冊上面自定義的日誌中間件
	r.Use(middleware.Logger())
	r.GET("/test", func(c *gin.Context) {
		// 查詢我們之前在日誌中間件，注入的鍵值數
		example := c.MustGet("example").(string)
		log.Println(example)
	})

	r.GET("/ping", func(c *gin.Context) {
		panic("errrrrrr")
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8888") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
