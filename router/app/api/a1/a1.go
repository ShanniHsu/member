package a1

import (
	"github.com/gin-gonic/gin"
	"member/router/app/controller"
	"member/router/repository"
	"member/router/service"
	"net/http"
)

// 模擬一些私人數據
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func Init(router *gin.Engine) {
	newRepo := repository.NewRepository()
	newUserService := service.NewUserService(newRepo)
	api := controller.NewAppController(newUserService)
	router.POST("/register", api.Register)
	router.POST("/login", api.Login)

	// BasicAuth
	// 路由組使用gin.BasicAuth() 中間件
	// gin.Accounts 是map[string]string的一種快捷方式
	authorized := router.Group("/user", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	authorized.GET("/secrets", func(c *gin.Context) {
		// 獲取用戶，它是由BasicAuth中間件設置的
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})
	// Postman Auth use "Basic Auth" & input Username and Password
	/*
		Ex1: if input Username is foo & Password is bar , the result is
		{
		   "secret": {
		       "email": "foo@bar.com",
		       "phone": "123433"
		   },
		   "user": "foo"
		}
	*/

	/*
			Ex2: if input Username is manu & Password is 4321 , the result is
		    {
		    "secret": "NO SECRET :(",
		    "user": "manu"
		    }
	*/
}
