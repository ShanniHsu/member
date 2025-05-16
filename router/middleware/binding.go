package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"member/pkg/message"
)

func Binding[T any]() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp := message.Response{Msg: "System error"}
		var obj T
		if err := ctx.ShouldBindJSON(&obj); err != nil {
			resp.Msg = "Invalid request body"
			resp.ResponseBadRequest(ctx)
			ctx.Abort()
			return
		}

		msg := "Please check your input"
		if err := validator.New().Struct(obj); err != nil {
			// 如果錯誤不是validator回傳的
			if _, ok := err.(validator.ValidationErrors); !ok {
				resp.Msg = "Validation failed"
				resp.ResponseBadRequest(ctx)
				log.Println("error:", err.Error())
				ctx.Abort()
				return
			}

			validatorErr := err.(validator.ValidationErrors)[0]
			switch validatorErr.Tag() {
			case "required":
				msg = validatorErr.Field() + " is required"
			case "max":
				msg = validatorErr.Field() + " is too long"
			case "min":
				msg = validatorErr.Field() + " is too short"
			case "email":
				msg = validatorErr.Field() + " is invalid format"
			}

			resp.Msg = msg
			resp.ResponseBadRequest(ctx)
			log.Println("error:", validatorErr.Error())
			ctx.Abort()
			return
		}

		//統一使用指標格式去儲存，要使用的時候要用指標格式取出
		ctx.Set("validatedBody", &obj)
		ctx.Next()
	}
}
