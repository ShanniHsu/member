package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cors.New(cors.Config{
			AllowOrigins:     []string{"https://foo.com"},
			AllowMethods:     []string{"PUT", "PATCH"},
			AllowHeaders:     []string{"Origin"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				return origin == "https://github.com"
			},
			MaxAge: 12 * time.Hour,
		})
		// AllowOrigins: 允許https://foo.com的跨域請求
		// AllowMethods: 僅允許使用PUT&PATCH方法的跨域請求
		// AllowHeaders: 跨域請求可以帶有Origin請求頭
		// ExposeHeaders: 響應中允許客戶端讀取Content-Length頭
		// AllowCredentials: 是否允許跨域攜帶憑據(如Cookies、Http認證信息等)
		// AllowOriginFunc: 如果來源是https://github.com，動態允許其跨域訪問
		// MaxAge: 瀏覽器在跨域請求預檢後，可以緩存此配置12小時
	}
}
