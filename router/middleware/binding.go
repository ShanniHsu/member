package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"member/pkg/message"
)

func Binding[T any]() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp := message.Response{MsgID: message.SYSTEM_ERROR}
		var obj T
		if err := ctx.ShouldBindJSON(&obj); err != nil {
			resp.MsgID = message.INVALID_REQUEST_BODY
			resp.ResponseBadRequest(ctx)
			log.Println("Invalid request body")
			ctx.Abort()
			return
		}

		// 是否符合設定的tag條件
		if err := validator.New().Struct(obj); err != nil {
			// 如果錯誤不是validator回傳的
			if _, ok := err.(validator.ValidationErrors); !ok {
				resp.MsgID = message.NON_VALIDATION_ERROR
				resp.ResponseBadRequest(ctx)
				log.Println("error:", err.Error())
				ctx.Abort()
				return
			}

			validatorErr := err.(validator.ValidationErrors)[0]

			field := map[string]interface{}{
				"field": validatorErr.Field(),
			}

			switch validatorErr.Tag() {
			case "required":
				resp.MsgID = message.FIELD_IS_REQUIRED
			case "max":
				resp.MsgID = message.FIELD_IS_TOO_LONG
			case "min":
				resp.MsgID = message.FIELD_IS_TOO_SHORT
			case "email":
				resp.MsgID = message.FIELD_IS_INVALID_EMAIL_FORMAT
			default:
				resp.MsgID = message.FIELD_CHECK_INPUT_PLEASE
			}

			resp.ResponseBadRequest(ctx, field)
			log.Println("error:", validatorErr.Error())
			ctx.Abort()
			return
		}

		//統一使用指標格式去儲存，要使用的時候要用指標格式取出
		ctx.Set("validatedBody", &obj)
		ctx.Next()
	}
}
