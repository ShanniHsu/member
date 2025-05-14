package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"member/models"
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
	var idInt, typeInt int64
	var err error

	userCtx, exist := ctx.Get("user")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized!",
		})
		log.Print("User in context is missing")
		return
	}

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

	req := &get_restaurants.Request{
		ID:   idInt,
		Type: typeInt,
	}

	user := userCtx.(*models.User)
	resp, err := c.restaurantService.GetRestaurantList(user, req)
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
	return
}
