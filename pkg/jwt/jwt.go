package jwt

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var privateKey *ecdsa.PrivateKey
var Issuer string

func InitJwt() {
	privateKey = new(ecdsa.PrivateKey)
	Issuer = "ithome"
}

type AuthClaims struct {
	jwt.RegisteredClaims
	Token string
}

// 獲取使用者資訊
func GetUserInfo(ctx *gin.Context) (claim interface{}, err error) {
	// 通過http header中的token解析來認證
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		return nil, fmt.Errorf("no jwt token")
	}

	claim, err = ParseToken(token[7:])
	if err != nil {
		return nil, fmt.Errorf("bad jwt: %s", err)
	}

	return claim, nil
}

// 解析jwt token
func ParseToken(token string) (authClaims interface{}, err error) {
	jwtToken, err := jwt.ParseWithClaims(token, &AuthClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return &privateKey.PublicKey, nil
	})

	if err != nil {
		return
	} else if jwtToken.Valid {
		authClaims = jwtToken.Claims
	} else {
		err = errors.New("ECDSA Token is invalid")
		return
	}
	return
}

// 產生jwt
func GenerateJWT(token string) (tokenString string, err error) {
	// 生成 ECDSA 私鑰(此處使用 P-256 橢圓曲線)
	privateKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return
	}
	expiresAt := time.Now().Add(10 * time.Minute)
	claims := AuthClaims{
		Token: token,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    Issuer,
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}
	// 建立Token，並設置簽名方法為ES256 (ECDSA SHA-256)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err = jwtToken.SignedString(privateKey)
	if err != nil {
		return
	}
	return
}
