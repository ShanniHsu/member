package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Recovery(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Server Error",
			})
		}
		ctx.Abort()
	}()
	ctx.Next()
}
