package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/config"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/auth"
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

			jwtExt, ok := claims["extend"].(map[string]interface{})
			if !ok {
				ginx.NewContext(c).Render(statusx.StatusVerifyTokenError, "nil", http.StatusUnauthorized)
				c.Abort()
				return
			}

			jwtExtRoleId, ok := jwtExt["role_id"].(float64)
			if !ok {
				ginx.NewContext(c).Render(statusx.StatusVerifyTokenError, "nil", http.StatusUnauthorized)
				c.Abort()
				return
			}
			jwtExtUId, ok := jwtExt["u_id"].(float64)
			if !ok {
				ginx.NewContext(c).Render(statusx.StatusVerifyTokenError, "nil", http.StatusUnauthorized)
				c.Abort()
				return
			}
			jwtExtMerchantId, ok := jwtExt["merchant_id"].(float64)
			if !ok {
				ginx.NewContext(c).Render(statusx.StatusVerifyTokenError, "nil", http.StatusUnauthorized)
				c.Abort()
				return
			}

			//校验用户信息
			if jwtExtUId <= 0 || jwtExtRoleId <= 0 {
				ginx.NewContext(c).Render(statusx.StatusVerifyTokenError, "nil", http.StatusUnauthorized)
				c.Abort()
				return
			}

			authInfo := auth.JwtExt{
				RoleId:     uint64(jwtExtRoleId),
				UId:        uint64(jwtExtUId),
				MerchantId: uint64(jwtExtMerchantId),
			}

			//TODO: 通过用户表的最近更新时间去判断是否需要更新权限，登出时间判断
			//TODO: 开发完成后打开此处逻辑
			//由于中间件执行逻辑，开始处理权限相关逻辑
			//perm := checkPermissions(c, strconv.FormatInt(int64(jwtExtRoleId), 10))
			//if !perm {
			//	ginx.NewContext(c).Render(statusx.StatusPermissionsError, "nil", http.StatusBadRequest)
			//	c.Abort()
			//	return
			//}

			c.Set(config.AUTHUSERKEY, authInfo)
			c.Request = c.Request.WithContext(c)
			c.Next()
		} else {
			ginx.NewContext(c).Render(statusx.StatusUnauthorized, "nil", http.StatusUnauthorized)
			c.Abort()
			return
		}
	}
}

func checkPermissions(c *gin.Context, roleId string) bool {
	//获取请求path
	obj := c.Request.URL.Path
	//获取请求方法
	act := c.Request.Method
	//获取用户的角色,从token解析出来
	sub := roleId

	e := auth.CasbinServiceApp.Casbin()

	//判断策略是否存在
	ok, err := e.Enforce(sub, obj, act)
	if err != nil {
		return false
	}

	return ok
}
