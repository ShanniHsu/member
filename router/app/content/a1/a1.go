package a1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func (r Response) ResponseOk(message string) {
	var ctx *gin.Context
	r.Message = message
	ctx.JSON(http.StatusOK, r)
}
