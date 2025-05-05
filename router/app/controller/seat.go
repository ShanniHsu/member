package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c appController) GetSeats(ctx *gin.Context) {
	err := c.seatService.GetSeats()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get all seats successfully",
	})
	return
}
