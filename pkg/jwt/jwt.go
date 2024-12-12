package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var Issuer string
var privateKey []byte
var publicKey []byte

func InitJwt() {
	privateKey = []byte("-----BEGIN EC PRIVATE KEY-----\nMHcCAQEEINDkpT6HR5JgRkdVII+ltXeW2cGbXbB8OS8c2lLPOWXqoAoGCCqGSM49\nAwEHoUQDQgAEOMx2uM9MGcu/e3VHhCgzCcrSArl9HATcwj5njs7ax6FXdofzuDlF\nx5lUgjWBJZc4pTfQ/IMKJLZFKvL8XMHPfg==\n-----END EC PRIVATE KEY-----")
	publicKey = []byte("-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEOMx2uM9MGcu/e3VHhCgzCcrSArl9\nHATcwj5njs7ax6FXdofzuDlFx5lUgjWBJZc4pTfQ/IMKJLZFKvL8XMHPfg==\n-----END PUBLIC KEY-----")
	Issuer = "ithome"
}

type AuthClaims struct {
	jwt.RegisteredClaims
	Token string
}

// 解析jwt token
func ParseToken(token string) (authClaims *AuthClaims, err error) {
	//jwtToken, err := jwt.ParseWithClaims(token, &AuthClaims{}, func(token *jwt.Token) (i interface{}, e error) {
	//	return &privateKey.PublicKey, nil
	//})
	pubKey, err := jwt.ParseECPublicKeyFromPEM(publicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}

	jwtToken, err := jwt.ParseWithClaims(token, &AuthClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return pubKey, nil
	})

	if err != nil {
		return
	} else if jwtToken.Valid {
		authClaims = jwtToken.Claims.(*AuthClaims)
	} else {
		err = errors.New("ECDSA Token is invalid")
		return
	}
	return
}

// 產生jwt
func GenerateJWT(token string) (tokenString string, err error) {
	// 生成 ECDSA 私鑰(此處使用 P-256 橢圓曲線)
	//privateKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	//if err != nil {
	//	return
	//}

	privKey, err := jwt.ParseECPrivateKeyFromPEM(privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %w", err)
	}

	expiresAt := time.Now().Add(10 * time.Hour)
	claims := AuthClaims{
		Token: token,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    Issuer,
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}
	// 建立Token，並設置簽名方法為ES256 (ECDSA SHA-256)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err = jwtToken.SignedString(privKey)
	if err != nil {
		return
	}
	return
}

// OpenSSL生成加密密鑰的常用工具，OpenSSL 生成 P-256 公私鑰對的步驟
// 生成私鑰
// openssl ecparam -genkey -name prime256v1 -noout -out ec_private.pem

// 生成公鑰
// openssl ec -in ec_private.pem -pubout -out ec_public.pem

// 查看私鑰
// cat ec_private.pem
// 查看公鑰
// cat ec_public.pem
