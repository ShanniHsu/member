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

		message := "Please check your input"
		if err := validator.New().Struct(obj); err != nil {
			// 如果錯誤不是validator回傳的
			if _, ok := err.(validator.ValidationErrors); !ok {
				ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
					"error":  "Validation failed",
					"detail": err.Error(),
				})
				return
			}

			validatorErr := err.(validator.ValidationErrors)[0]
			switch validatorErr.Tag() {
			case "required":
				message = validatorErr.Field() + " is required"
			case "max":
				message = validatorErr.Field() + " is too long"
			case "min":
				message = validatorErr.Field() + " is too short"
			case "email":
				message = validatorErr.Field() + " is invalid format"
			}

			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":  message,
				"detail": validatorErr.Error(),
			})
			return
		}

		//統一使用指標格式去儲存，要使用的時候要用指標格式取出
		ctx.Set("validatedBody", &obj)
		ctx.Next()
	}
}
