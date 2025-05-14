package a1

import (
	"github.com/gin-gonic/gin"
	create_user_restaurant "member/router/app/content/create-user-restaurant"
	delete_user_restaurant "member/router/app/content/delete-user-restaurant"
	"member/router/app/content/login"
	"member/router/app/content/register"
	"member/router/app/controller"
	"member/router/app/middleware"
	middleware2 "member/router/middleware"
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
	router.POST("/register", middleware2.Binding[register.Request](), api.Register) // 註冊
	router.POST("/login", middleware2.Binding[login.Request](), api.Login)          // 登入

	auth := router.Group("/auth")
	auth.Use(middleware.Auth(newUserService))
	auth.GET("/info", api.GetUserInfo)                                                                                   // 獲取個人資料
	auth.GET("/restaurants", api.GetRestaurants)                                                                         // 獲取餐廳列表
	auth.GET("/restaurant-list", api.GetRestaurantList)                                                                  // 獲取餐廳列表
	auth.GET("/pocket-restaurant-list", api.GetPocketRestaurantList)                                                     // 口袋餐廳列表
	auth.POST("/pocket-restaurant", middleware2.Binding[create_user_restaurant.Request](), api.AddPocketRestaurant)      // 加入口袋餐廳
	auth.DELETE("/pocket-restaurant", middleware2.Binding[delete_user_restaurant.Request](), api.DeletePocketRestaurant) // 移除口袋餐廳
	auth.POST("/logout", api.Logout)                                                                                     // 登出
}
