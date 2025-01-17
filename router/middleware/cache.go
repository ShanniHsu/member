package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
)

var cache sync.Map // 全局緩存

func Cache(ctx *gin.Context) {
	// 使用請求路徑作為緩存鍵
	cacheKey := ctx.Request.URL.Path

	// 檢查緩存
	if value, ok := cache.Load(cacheKey); ok {
		// 如果有緩存，直接返回
		ctx.JSON(http.StatusOK, gin.H{
			"from_cache": true,
			"data":       value,
		})
		// 中止後續處理
		ctx.Abort()
		return
	}
	ctx.Next()
	// 緩存響應結果
	cache.Store(cacheKey, gin.H{
		"response": "This is a cached response!",
		"time":     time.Now(),
	})
}
