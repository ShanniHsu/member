package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"member/models"
	"member/pkg/message"
	create_user_restaurant "member/router/app/content/create-user-restaurant"
	delete_user_restaurant "member/router/app/content/delete-user-restaurant"
	get_user_restaurants "member/router/app/content/get-user-restaurants"
	"strconv"
)

func (c appController) GetPocketRestaurantList(ctx *gin.Context) {
	resp := message.Response{MsgID: message.SYSTEM_ERROR}
	userCtx, exist := ctx.Get("user")
	if !exist {
		resp.MsgID = message.UNAUTHORIZED
		resp.ResponseUnauthorized(ctx)
		log.Print("User in context is missing")
		return
	}

	var idInt, typeInt int64
	var err error

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

	name := ctx.Query("name")

	address := ctx.Query("address")

	page := ctx.Query("page")
	if page == "" {
		page = "1"
	} else {
		_, err = strconv.Atoi(page)
		if err != nil {
			resp.MsgID = message.FIELD_FORMAT_ERROR
			resp.ResponseBadRequest(ctx)
			return
		}
	}

	pageSize := ctx.Query("page_size")
	if pageSize == "" {
		pageSize = "100"
	} else {
		_, err = strconv.Atoi(pageSize)
		if err != nil {
			resp.MsgID = message.FIELD_FORMAT_ERROR
			resp.ResponseBadRequest(ctx)
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
		resp.MsgID = err.Error()
		resp.ResponseBadRequest(ctx)
		return
	}

	resp.MsgID = message.GET_POCKET_RESTAURANT_LIST_SUCCESSFULLY
	resp.Data = data
	resp.ResponseSuccess(ctx)
	return
}

func (c appController) AddPocketRestaurant(ctx *gin.Context) {
	resp := message.Response{MsgID: message.SYSTEM_ERROR}
	userCtx, exist := ctx.Get("user")
	if !exist {
		resp.MsgID = message.UNAUTHORIZED
		resp.ResponseUnauthorized(ctx)
		log.Println("User in context is missing")
		return
	}

	value, exist := ctx.Get("validatedBody")
	if !exist {
		resp.MsgID = message.VALIDATED_DATA_MISSING
		resp.ResponseBadRequest(ctx)
		return
	}

	user := userCtx.(*models.User)
	req := value.(*create_user_restaurant.Request)
	err := c.userRestaurant.AddPocketRestaurant(user, req)
	if err != nil {
		resp.MsgID = err.Error()
		resp.ResponseBadRequest(ctx)
		return
	}

	resp.MsgID = message.ADD_POCKET_RESTAURANT_SUCESSFULLY
	resp.ResponseSuccess(ctx)
	return
}

func (c appController) DeletePocketRestaurant(ctx *gin.Context) {
	resp := message.Response{MsgID: message.SYSTEM_ERROR}
	userCtx, exist := ctx.Get("user")
	if !exist {
		resp.MsgID = message.UNAUTHORIZED
		resp.ResponseUnauthorized(ctx)
		log.Println("User in context is missing")
		return
	}

	value, exist := ctx.Get("validatedBody")
	if !exist {
		resp.MsgID = message.VALIDATED_DATA_MISSING
		resp.ResponseBadRequest(ctx)
		return
	}

	user := userCtx.(*models.User)
	req := value.(*delete_user_restaurant.Request)
	err := c.userRestaurant.DeletePocketRestaurant(user, req)
	if err != nil {
		resp.MsgID = err.Error()
		resp.ResponseBadRequest(ctx)
		return
	}

	resp.MsgID = message.DELETE_POCKET_RESTAURANT_SUCCESSFULLY
	resp.ResponseSuccess(ctx)
	return
}
