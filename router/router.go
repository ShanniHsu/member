package router

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
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
	middleware.AddGinMiddleware(router, "requestId", "cors", "logger")
	var s = securecookie.New([]byte("very-secret-key"), nil)

	router.GET("/set", func(c *gin.Context) {
		// 設置加密的 Cookie
		encoded, err := s.Encode("secureCookie", "SensitiveData")
		if err != nil {
			log.Println("Error encoding cookie:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to set cookie"})
			return
		}
		http.SetCookie(c.Writer, &http.Cookie{
			Name:     "secureCookie",
			Value:    encoded,
			Path:     "/",
			HttpOnly: true,
		})
		c.JSON(http.StatusOK, gin.H{"message": "Cookie set"})
	})

	router.GET("/get", func(c *gin.Context) {
		// 獲取並解密 Cookie
		cookie, err := c.Request.Cookie("secureCookie")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "no cookie found"})
			return
		}
		var value string
		err = s.Decode("secureCookie", cookie.Value, &value)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to decode cookie"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"decryptedValue": value})
	})

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
