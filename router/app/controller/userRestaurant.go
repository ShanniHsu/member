package controller

import (
	"github.com/gin-gonic/gin"
	create_user_restaurant "member/router/app/content/create-user-restaurant"
	"net/http"
)

func (c appController) AddPocketRestaurant(ctx *gin.Context) {
	req := new(create_user_restaurant.Request)
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = c.userRestaurant.AddPocketRestaurant(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Add pocket restaurant successfully!",
	})
	return
}
