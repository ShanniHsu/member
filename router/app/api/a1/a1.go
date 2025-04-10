package a1

import (
	"github.com/gin-gonic/gin"
	"member/router/app/controller"
	"member/router/app/middleware"
	"member/router/repository"
	"member/router/service"
)

func Init(router *gin.Engine) {
	newRepo := repository.NewRepository()
	newUserService := service.NewUserService(newRepo)
	newRestaurantService := service.NewRestaurantService(newRepo)
	newUserRestaurantService := service.NewUserRestaurantService(newRepo)
	api := controller.NewAppController(
		newUserService,
		newRestaurantService,
		newUserRestaurantService,
	)
	router.POST("/register", api.Register) // 註冊
	router.POST("/login", api.Login)       // 登入

	auth := router.Group("/auth")
	auth.Use(middleware.Auth(newUserService))
	auth.GET("/info", api.GetUserInfo)                               // 獲取個人資料
	auth.GET("/restaurants", api.GetRestaurants)                     // 獲取餐廳列表
	auth.GET("/restaurant-list", api.GetRestaurantList)              // 獲取餐廳列表
	auth.GET("/pocket-restaurant-list", api.GetPocketRestaurantList) // 口袋餐廳列表
	auth.POST("/pocket-restaurant", api.AddPocketRestaurant)         // 加入口袋餐廳
	auth.DELETE("/pocket-restaurant", api.DeletePocketRestaurant)    // 移除口袋餐廳
	auth.POST("/logout", api.Logout)                                 // 登出
}
