package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/docs"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/bootstrap"
)

var (
	routerNoCheckRole = make([]func(rg *gin.RouterGroup), 0)
	routerCheckRole   = make([]func(rg *gin.RouterGroup, gHanFun ...gin.HandlerFunc), 0)
)

func ExecRouter(r *gin.Engine, gHanFun ...gin.HandlerFunc) *gin.Engine {
	if bootstrap.IsDevelopment() {
		//先注册swagger组件
		registerSwagger(r)
	}

	// 无需认证的路由
	noCheckRoleRouter(r)
	// 需要认证的路由
	checkRoleRouter(r, gHanFun...)

	return r
}

// 无需认证的路由示例
func noCheckRoleRouter(r *gin.Engine) {
	// 可根据业务需求来设置接口版本
	rg := r.Group("/merchant/api/v1")
	for _, f := range routerNoCheckRole {
		f(rg)
	}
}

// 需要认证的路由示例
func checkRoleRouter(r *gin.Engine, gHanFun ...gin.HandlerFunc) {
	// 可根据业务需求来设置接口版本
	rg := r.Group("/merchant/api/v1")
	for _, f := range routerCheckRole {
		f(rg, gHanFun...)
	}
}

func registerSwagger(r gin.IRouter) {
	// API文档访问地址: http://host/swagger/index.html
	// 注解定义可参考 https://github.com/swaggo/swag#declarative-comments-format
	// 样例 https://github.com/swaggo/swag/blob/master/example/basic/api/api.go
	docs.SwaggerInfo.BasePath = "/merchant/api/v1"
	docs.SwaggerInfo.Title = "商户管理后台"
	docs.SwaggerInfo.Description = "实现一个管理系统的后端API服务"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8099"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
