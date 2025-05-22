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
	resp := message.Response{MsgID: message.SYSTEM_ERROR}
	data, err := c.restaurantService.GetRestaurants()
	if err != nil {
		resp.MsgID = err.Error()
		resp.ResponseBadRequest(ctx)
		return
	}

	resp.MsgID = message.GET_RESTAURANT_LIST_SUCCESSFULLY
	resp.Data = data
	resp.ResponseSuccess(ctx)
	return
}

func (c appController) GetRestaurantList(ctx *gin.Context) {
	var idInt, typeInt int64
	var err error

	resp := message.Response{MsgID: message.SYSTEM_ERROR}
	userCtx, exist := ctx.Get("user")
	if !exist {
		resp.MsgID = message.UNAUTHORIZED
		resp.ResponseUnauthorized(ctx)
		log.Println("User in context is missing")
		return
	}

	idString := ctx.Query("id")
	if idString != "" {
		idInt, err = strconv.ParseInt(idString, 10, 64)
		if err != nil {
			resp.MsgID = message.FIELD_FORMAT_ERROR
			resp.ResponseBadRequest(ctx)
			return
		}
	}

	typeString := ctx.Query("type")

	if typeString != "" {
		typeInt, err = strconv.ParseInt(typeString, 10, 64)
		if err != nil {
			resp.MsgID = message.FIELD_FORMAT_ERROR
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
		resp.MsgID = err.Error()
		resp.ResponseBadRequest(ctx)
		return
	}

	resp.MsgID = message.GET_RESTAURANT_LIST_SUCCESSFULLY
	resp.Data = data
	resp.ResponseSuccess(ctx)
	return
}
