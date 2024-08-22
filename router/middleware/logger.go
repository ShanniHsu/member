package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// 可以通過上下文對象，設置一些依附在上下文對象裡面的鍵/值數據
		c.Set("example", "12345")

		// 在這裡處理請求到達控制器函數之前的邏輯

		// 調用下一個中間件，或者控制器處理函數，具體得看註冊多少個中間件
		c.Next()

		// 在這裡可以處理請求返回給用戶之前的邏輯
		latency := time.Since(t)
		log.Print(latency)

		// 例如，查詢請求狀態碼
		status := c.Writer.Status()
		log.Println(status)
	}
}
