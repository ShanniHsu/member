package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Result  bool        `json:"result"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r Response) ResponseSuccess(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"result": r.Result, "message": r.Message, "data": r.Data})
	return
}

func (r Response) ResponseServerError(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{"result": r.Result, "message": r.Message})
}

func (r Response) ResponseBadRequest(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, gin.H{"result": r.Result, "message": r.Message})
}
