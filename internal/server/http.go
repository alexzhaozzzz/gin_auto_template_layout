package server

import (
	"github.com/gin-gonic/gin"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/router"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/router/middleware"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/bootstrap"
)

func InitServer() *gin.Engine {
	r := gin.New()
	if r == nil {
		panic("load gin error")
	}

	if bootstrap.IsDevelopment() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// 自定义错误处理
	r.Use(middleware.CustomError)
	// NoCache is a middleware function that appends headers
	//r.Use(middleware.NoCache)
	// 跨域处理
	r.Use(middleware.Options)

	// the jwt middleware
	authMid := middleware.JWTAuth()
	startAt := middleware.StartAt()

	// 注册业务路由
	router.ExecRouter(r, startAt, authMid)

	return r
}
