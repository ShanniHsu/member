package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Timing(ctx *gin.Context) {
	// 記錄時間
	start := time.Now()
	// 執行下一個處理器
	ctx.Next()
	// 記錄請求消耗時間
	duration := time.Since(start)
	// 記錄日誌
	log.Printf("Request %s %s took %v", ctx.Request.Method, ctx.Request.URL.Path, duration)
}
