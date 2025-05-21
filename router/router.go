package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"member/router/app/api/a1"
	"member/router/middleware"
	"net/http"
)

func Init() {
	r := newRouter()
	addr := viper.GetString("web.addr")
	port := viper.GetString("web.port")
	routerAddr := fmt.Sprintf("%s:%s", addr, port)
	r.Run(routerAddr) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func newRouter() *gin.Engine {
	router := gin.New()
	// 註冊上面自定義的日誌中間件
	middleware.AddGinMiddleware(router, "requestId", "cors", "logger", "locale")
	router.GET("/test", func(c *gin.Context) {
		// 查詢我們之前在日誌中間件，注入的鍵值數
		example := c.MustGet("example").(string)
		log.Println(example)
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	a1.Init(router)
	return router
}
