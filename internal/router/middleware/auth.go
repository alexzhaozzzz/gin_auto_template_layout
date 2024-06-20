package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/config"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/ginx"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/statusx"
	"net/http"
	"time"
)

// JWTAuth jwt验证
func JWTAuth() gin.HandlerFunc {
	secret := viper.GetString("jwt.secret")
	secretKey := []byte(secret) // 替换为你的256位密钥
	return func(c *gin.Context) {
		// 从请求头中获取token字符串
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			ginx.NewContext(c).Render(statusx.StatusUnauthorized, "nil", http.StatusUnauthorized)
			c.Abort()
			return
		}
		// 去掉Bearer前缀
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 验证alg是否为HS256
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secretKey, nil // 返回用于验证的密钥
		})
		if err != nil {
			ginx.NewContext(c).Render(statusx.StatusUnauthorized, "nil", http.StatusUnauthorized)
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			expDate, err := claims.GetExpirationTime()
			if err != nil {
				ginx.NewContext(c).Render(statusx.StatusVerifyTokenError, "nil", http.StatusUnauthorized)
				c.Abort()
				return
			}

			// 检查 token 是否已经过期
			if expDate.Unix() < time.Now().Unix() {
				ginx.NewContext(c).Render(statusx.StatusVerifyTokenError, "nil", http.StatusUnauthorized)
				c.Abort()
				return
			}

			//TODO：校验一下用户信息

			c.Set(config.AUTHUSERKEY, claims["extend"])
			c.Request = c.Request.WithContext(c)
			c.Next()
		} else {
			ginx.NewContext(c).Render(statusx.StatusUnauthorized, "nil", http.StatusUnauthorized)
			c.Abort()
			return
		}
	}
}
