package jwt

import (
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
	ID      int
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
		}
	}
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
func GenerateJWT(id int, account string) (string, error) {
	expirseAt := time.Now().Add(10 * time.Second)

	claims := AuthClaims{
		ID:      id,
		Account: account,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    Issuer,
			ExpiresAt: jwt.NewNumericDate(expirseAt),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString([]byte(Secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
