package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"jzsg.com/mca/core/common"
)

const (
	issuer = "jzsg.com"
	expire = 10 * time.Hour
	key    = "jzsg"
)

// TokenAuthMiddleware check token
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//if c.ClientIP() == "127.0.0.1" || c.ClientIP() == "::1" {
		//	c.Next()
		//	return
		//}
		token := c.GetHeader("token")
		if token == "" {
			common.Response(c, errors.New("token is nil"), common.TokenNilErr, nil)
			c.Abort()
			return
		}
		t, err := ParseToken(token)
		if err != nil {
			common.Response(c, err, common.TokenInvalidErr, nil)
			c.Abort()
			return
		}
		c.Request.Header.Set("id", t)
		c.Next()
	}
}

// GenerateToken 生成jwt signed token
func GenerateToken(id string) string {
	//jwfConf := config.JWTInfo()
	claim := jwt.StandardClaims{
		//Audience:
		ExpiresAt: time.Now().Add(expire).Unix(),
		Id:        id,
		Issuer:    issuer,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	ss, err := token.SignedString([]byte(key))
	if err != nil {
		panic(err)
	}
	return ss
}

//ParseToken parse jwt token
func ParseToken(ss string) (string, error) {
	claims := jwt.StandardClaims{}
	token, err := jwt.ParseWithClaims(ss, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return "", err
	}
	if !token.Valid {
		return "", errors.New("invalid token")
	}
	return claims.Id, nil
}
