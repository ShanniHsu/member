package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func LocaleMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		locale, exist := ctx.Get("locale")
		if !exist {
			// 先查看使用者偏好的語言清單與順序優先
			// "zh-TW,zh;q=0.9,en-US;q=0.8,en;q=0.7,ja;q=0.6"
			// 代表zh-TW > zh > en-US > en > ja
			lang := ctx.GetHeader("Accept-Language")
			lang = strings.Split(lang, ",")[0]
			switch lang {
			case "zh-TW":
				locale = "zh"
			case "zh-CN":
				locale = "zh"
			case "en-US":
				locale = "en"
			default:
				locale = "en"
			}
		}
		ctx.Set("locale", locale)
		ctx.Next()
	}
}
