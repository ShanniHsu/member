package middleware

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func AddGinMiddleware(router gin.IRouter, args ...string) gin.IRoutes {
	if len(args) > 0 {
		for i := 0; i < len(args); i++ {
			switch args[i] {
			case "requestId":
				router.Use(requestid.New())
			case "cors":
				router.Use(Cors())
			case "logger":
				router.Use(Logger())
			case "cache":
				router.Use(Cache)
			}
		}
	}
	return router
}
