package jwt

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"time"
)

var Secret = os.Getenv("SECRET")
var Issuer = os.Getenv("Issuer")

type AuthClaims struct {
	jwt.RegisteredClaims
	Account string
	ID      int64
}

// Decode jwt
func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 通過http header中的token解析認證
		token := ctx.Request.Header.Get("token")
		if token == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "not jwt token!",
			})
			// 終止後續的handler繼續執行
			ctx.Abort()
			return
		}

		// 解析jwt是否正確，如果不正確則提前結束，正確就繼續
		_, err := parseToken(token)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Token is invalid or expired!",
			})
			ctx.Abort()
			return
		}
	}
}

// 獲取使用者資訊
func GetUserInfo(ctx *gin.Context) (*AuthClaims, error) {
	// 通過http header中的token解析來認證
	token := ctx.Request.Header.Get("token")
	if token == "" {
		return nil, fmt.Errorf("no jwt token")
	}

	claim, err := parseToken(token)
	if err != nil {
		return nil, fmt.Errorf("bad jwt: %s", err)
	}

	return claim, nil
}

// 解析jwt token
func parseToken(token string) (*AuthClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &AuthClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(Secret), nil
	})

	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*AuthClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err
}

// 產生jwt
func GenerateJWT(id int64, account string) (tokenString string, err error) {
	expiresAt := time.Now().Add(10 * time.Second)

	claims := AuthClaims{
		ID:      id,
		Account: account,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    Issuer,
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err = token.SignedString([]byte(Secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
