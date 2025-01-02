package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 設置 CORS 標頭
		ctx.Header("Access-Control-Allow-Origin", "http://localhost:8080")            // 允許的前端地址
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // 允許的 HTTP 方法
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")     // 允許的請求頭
		ctx.Header("Access-Control-Allow-Credentials", "true")                        // 是否允許攜帶憑據（如 Cookie）

		// 如果是 OPTIONS 預檢請求，直接返回狀態 204
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}

		// 繼續處理其他請求
		ctx.Next()

		// https://cloud.tencent.com/developer/article/2378435
		// 允許所有來源
		// cors.Default()

		// 但這個沒有允許所有來源，還是必須使用AllowOrigins
		//config := cors.DefaultConfig()
		//config.AllowOrigins = []string{"https://foo.com"}

		/*
			curl -X OPTIONS http://localhost:8888/ping \
			     -H "Origin: https://foo.com" \
			     -H "Access-Control-Request-Method: PUT" \
			     -H "Access-Control-Request-Headers: Origin,Content-Type" \
			     -v
		*/
		// 這邊是比較嚴謹的用法
		//cors.New(cors.Config{
		//	AllowOrigins: []string{"http://localhost:8080"},
		//	AllowMethods: []string{"GET, POST, PUT, DELETE, OPTIONS"},
		//	AllowHeaders: []string{"Content-Type, Authorization"},
		//ExposeHeaders:    []string{"Content-Length"},
		//AllowCredentials: true,
		//AllowOriginFunc: func(origin string) bool {
		//	return origin == "https://github.com"
		//},
		//MaxAge: 1 * time.Hour,
		//})

		//method := ctx.Request.Method

		//放行所有OPTIONS方法
		//if method == "OPTIONS" {
		//	ctx.AbortWithStatus(http.StatusNoContent)
		//}
		// 這邊要解決跨域問題是先發一次options請求，獲取allowheader，允許跨域之後才會再發真正的Post請求

		//處理請求
		//ctx.Next()
		// AllowOrigins: 允許https://foo.com的跨域請求
		// AllowMethods: 僅允許使用PUT&PATCH方法的跨域請求
		// AllowHeaders: 跨域請求可以帶有Origin請求頭
		// ExposeHeaders: 響應中允許客戶端讀取Content-Length頭
		// AllowCredentials: 是否允許跨域攜帶憑據(如Cookies、Http認證信息等)
		// AllowOriginFunc: 如果來源是https://github.com，動態允許其跨域訪問
		// MaxAge: 瀏覽器在跨域請求預檢後，可以緩存此配置12小時
	}
}
