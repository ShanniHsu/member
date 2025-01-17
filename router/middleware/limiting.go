package middleware

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
)

// Define a rate limiter with 1 request per second and a burst of 5.
var limiter = rate.NewLimiter(1, 5)

// Middleware to check the rate limit.
func RateLimiter(ctx *gin.Context) {
	if !limiter.Allow() {
		ctx.JSON(http.StatusTooManyRequests, gin.H{"error": "too many request"})
		ctx.Abort()
		return
	}
	ctx.Next()
}
