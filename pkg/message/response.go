package message

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 補充: omitempty是將未賦值的字段省略(0,"")
// 更深入研究就回頭看 https://cloud.tencent.com/developer/article/2224409
type Response struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func (r Response) ResponseSuccess(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, r)
	return
}

func (r Response) ResponseBadRequest(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, r)
	return
}

func (r Response) ResponseUnauthorized(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, r)
	return
}
