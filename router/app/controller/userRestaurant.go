package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"member/models"
	create_user_restaurant "member/router/app/content/create-user-restaurant"
	delete_user_restaurant "member/router/app/content/delete-user-restaurant"
	get_user_restaurants "member/router/app/content/get-user-restaurants"
	"net/http"
	"strconv"
)

func (c appController) GetPocketRestaurantList(ctx *gin.Context) {
	userCtx, exist := ctx.Get("user")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized!",
		})
		log.Print("User in context is missing")
		return
	}

	var idInt, typeInt int64
	var err error

	idString := ctx.Query("id")
	if idString != "" {
		idInt, err = strconv.ParseInt(idString, 10, 64)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
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

	name := ctx.Query("name")

	address := ctx.Query("address")

	page := ctx.Query("page")
	if page == "" {
		page = "1"
	} else {
		_, err := strconv.Atoi(page)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	pageSize := ctx.Query("page_size")
	if pageSize == "" {
		pageSize = "100"
	} else {
		_, err := strconv.Atoi(pageSize)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	req := &get_user_restaurants.Request{
		ID:       idInt,
		Type:     typeInt,
		Name:     name,
		Address:  address,
		Page:     page,
		PageSize: pageSize,
	}

	user := userCtx.(*models.User)
	data, err := c.userRestaurant.GetPocketRestaurantList(user, req)
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
	userCtx, exist := ctx.Get("user")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized!",
		})
		log.Println("User in context is missing")
		return
	}

	value, exist := ctx.Get("validatedBody")
	if !exist {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Validated data missing",
		})
		return
	}

	user := userCtx.(*models.User)
	req := value.(*create_user_restaurant.Request)
	err := c.userRestaurant.AddPocketRestaurant(user, req)
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
	userCtx, exist := ctx.Get("user")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized!",
		})
		log.Println("User in context is missing")
		return
	}

	value, exist := ctx.Get("validatedBody")
	if !exist {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Validated data missing",
		})
		return
	}

	user := userCtx.(*models.User)
	req := value.(*delete_user_restaurant.Request)
	err := c.userRestaurant.DeletePocketRestaurant(user, req)
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
