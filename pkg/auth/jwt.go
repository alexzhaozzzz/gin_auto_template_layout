package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/config"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/ginx"
)

type JwtExt struct {
	UId        uint64 `json:"u_id"`
	MerchantId uint64 `json:"merchant_id"`
}

func NewJwt(secret, sub, name string, expTime int32, ext *JwtExt) string {
	// 定义密钥，确保这个密钥在验证JWT时也能访问到
	var signingKey = []byte(secret)

	// 定义一个 MapClaims 对象，存放自定义的数据声明
	claims := jwt.MapClaims{
		"sub":    sub,                                                       // 主题（Subject）
		"name":   name,                                                      // 名称
		"iat":    time.Now().Unix(),                                         // 签发时间
		"exp":    time.Now().Add(time.Hour * time.Duration(expTime)).Unix(), // 过期时间，这里设置为24小时后
		"extend": ext,                                                       // 扩展内容
	}

	// 创建一个 Token 对象，使用 HMAC 算法 HS256 进行签名
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥对Token进行签名
	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return ""
	}

	return tokenString
}

// GetJwtExt ...
func GetJwtExt(ctx *ginx.Context) (*JwtExt, bool) {
	ext, ok := ctx.Get(config.AUTHUSERKEY)
	if !ok {
		return nil, false
	}

	extInfo, ok := ext.(JwtExt)
	if !ok {
		return nil, false
	}

	return &extInfo, true
}
