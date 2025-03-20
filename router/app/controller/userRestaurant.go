package controller

import (
	"github.com/gin-gonic/gin"
	create_user_restaurant "member/router/app/content/create-user-restaurant"
	delete_user_restaurant "member/router/app/content/delete-user-restaurant"
	get_user_restaurants "member/router/app/content/get-user-restaurants"
	"net/http"
	"strconv"
)

func (c appController) GetPocketRestaurantList(ctx *gin.Context) {
	req := new(get_user_restaurants.Request)

	idString := ctx.Query("id")
	if idString != "" {
		idInt, err := strconv.ParseInt(idString, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
			})
			return
		}
		req.ID = idInt
	}

	typeString := ctx.Query("type")
	if typeString != "" {
		typeInt, err := strconv.ParseInt(typeString, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		req.Type = typeInt
	}

	name := ctx.Query("name")
	if name != "" {
		req.Name = name
	}

	address := ctx.Query("address")
	if address != "" {
		req.Address = address
	}

	page := ctx.Query("page")
	if page != "" {
		_, err := strconv.Atoi(page)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		req.Page = page
	}

	if req.Page == "" {
		req.Page = "1"
	}

	pageSize := ctx.Query("page_size")
	if pageSize != "" {
		_, err := strconv.Atoi(pageSize)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		req.PageSize = pageSize
	}

	if req.PageSize == "" {
		req.PageSize = "100"
	}

	data, err := c.userRestaurant.GetPocketRestaurantList(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get pocket restaurant list!",
		"data":    data,
	})
	return
}

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

func (c appController) DeletePocketRestaurant(ctx *gin.Context) {
	req := new(delete_user_restaurant.Request)
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = c.userRestaurant.DeletePocketRestaurant(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete pocket restaurant successfully!",
	})
	return
}
