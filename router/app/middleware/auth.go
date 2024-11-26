package middleware

import (
	"github.com/gin-gonic/gin"
	"member/models"
	"member/pkg/jwt"
	"member/router/service"
	"net/http"
	"strings"
)

// 路由請求中間件，前端必須把token放到請求頭上，對服務器進行驗證token成功後，才能訪問後續的請求路由
func Auth(userService service.User) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user = new(models.User)

		// 獲取authorization header: 獲取前端傳過來的訊息
		tokenString := ctx.GetHeader("Authorization")

		// 驗證前端傳過來的token格式，須不為空，且開頭為Bearer
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Request failed!",
			})
			ctx.Abort()
			return
		}

		// 驗證通過
		token := tokenString[7:]
		claims, err := jwt.ParseToken(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token is invalid or expired!",
			})
			ctx.Abort()
			return
		}

		user, err = userService.AuthBearerToken(claims.Token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token is invalid or expired!",
			})
			ctx.Abort()
			return
		}

		ctx.Set("user", user)
		ctx.Next()
		return
	}
}
