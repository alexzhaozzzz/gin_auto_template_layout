package router

import (
	"github.com/gin-gonic/gin"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/ginx"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, noCheckProRouter)
}

func noCheckProRouter(r *gin.RouterGroup) {
	v1 := r.Group("/")
	{
		template := logic.Template{}
		user := logic.NewUser()

		v1.POST("/t_add", ginx.Handle(template.Add))
		v1.POST("/test", ginx.Handle(template.Test))

		v1.POST("/login", ginx.Handle(user.Login))
	}
}
