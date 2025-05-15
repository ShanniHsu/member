package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"member/models"
	"member/pkg/message"
	get_restaurants "member/router/app/content/get-restaurants"
	"strconv"
)

func (c appController) GetRestaurants(ctx *gin.Context) {
	resp := message.Response{Msg: "System Error"}
	data, err := c.restaurantService.GetRestaurants()
	if err != nil {
		resp.Msg = err.Error()
		resp.ResponseBadRequest(ctx)
		return
	}

	resp.Msg = "Get restaurants successfully!"
	resp.Data = data
	resp.ResponseSuccess(ctx)
	return
}

func (c appController) GetRestaurantList(ctx *gin.Context) {
	var idInt, typeInt int64
	var err error

	resp := message.Response{Msg: "System Error"}
	userCtx, exist := ctx.Get("user")
	if !exist {
		resp.Msg = "Unauthorized!"
		resp.ResponseUnauthorized(ctx)
		log.Println("User in context is missing")
		return
	}

	idString := ctx.Query("id")
	if idString != "" {
		idInt, err = strconv.ParseInt(idString, 10, 64)
		if err != nil {
			resp.Msg = err.Error()
			resp.ResponseBadRequest(ctx)
			return
		}
	}

	typeString := ctx.Query("type")

	if typeString != "" {
		typeInt, err = strconv.ParseInt(typeString, 10, 64)
		if err != nil {
			resp.Msg = err.Error()
			resp.ResponseBadRequest(ctx)
			return
		}
	}

	req := &get_restaurants.Request{
		ID:   idInt,
		Type: typeInt,
	}

	user := userCtx.(*models.User)
	data, err := c.restaurantService.GetRestaurantList(user, req)
	if err != nil {
		resp.Msg = err.Error()
		resp.ResponseBadRequest(ctx)
		return
	}

	resp.Msg = "Get restaurant list successfully!"
	resp.Data = data
	resp.ResponseSuccess(ctx)
	return
}
