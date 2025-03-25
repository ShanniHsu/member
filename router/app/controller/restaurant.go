package controller

import (
	"github.com/gin-gonic/gin"
	get_restaurants "member/router/app/content/get-restaurants"
	"net/http"
	"strconv"
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

func (c appController) GetRestaurantList(ctx *gin.Context) {
	req := new(get_restaurants.Request)
	var idInt, typeInt int64
	var err error

	idString := ctx.Query("id")
	if idString != "" {
		idInt, err = strconv.ParseInt(idString, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	typeString := ctx.Query("type")

	if typeString != "" {
		typeInt, err = strconv.ParseInt(typeString, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	req.Type = typeInt
	req.ID = idInt

	resp, err := c.restaurantService.GetRestaurantList(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get restaurant list successfully!",
		"data":    resp,
	})
}
