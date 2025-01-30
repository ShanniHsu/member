package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c appController) GetRestaurants(ctx *gin.Context) {
	resp, err := c.restaurantService.GetRestaurants()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get restaurants successfully!",
		"data":    resp,
	})
}
