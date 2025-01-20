package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

var cache sync.Map // 全局緩存

// 自定義ResponseWriter，補獲響應Body
// 因為大部分框架不直接提供響應補獲功能，但都支持通過自定義的方式實現
// 基本原則是包裝或攔截響應處理流程
// 捕獲響應在需要進行日誌記錄、性能監控或響應緩存的場景下非常有用，但需要考慮性能影響
type ResponseRecorder struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (rw *ResponseRecorder) Write(b []byte) (int, error) {
	rw.body.Write(b) //寫入緩存
	return rw.ResponseWriter.Write(b)
}

func Cache(ctx *gin.Context) {
	// 使用請求路徑作為緩存鍵
	cacheKey := ctx.Request.URL.Path

	// 檢查緩存
	if value, ok := cache.Load(cacheKey); ok {
		var cacheResponse map[string]interface{}
		if err := json.Unmarshal([]byte(value.(string)), &cacheResponse); err == nil {
			// 緩存命中
			ctx.JSON(http.StatusOK, cacheResponse)
			// 中止後續處理
			ctx.Abort()
			return
		}
	}
	recorder := &ResponseRecorder{
		ResponseWriter: ctx.Writer,
		body:           bytes.NewBufferString(""),
	}
	ctx.Writer = recorder

	ctx.Next()

	status := ctx.Writer.Status()
	if status == http.StatusOK {
		cache.Store(cacheKey, recorder.body.String())
	}
}
