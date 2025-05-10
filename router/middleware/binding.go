package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func Binding[T any]() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var obj T
		if err := ctx.ShouldBindJSON(&obj); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":  "Invalid request body",
				"detail": err.Error(),
			})
			return
		}

		if err := validator.New().Struct(obj); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":  "Validation failed",
				"detail": err.Error(),
			})
			return
		}

		//統一使用指標格式去儲存，要使用的時候要用指標格式取出
		ctx.Set("validatedBody", &obj)
		ctx.Next()
	}
}
