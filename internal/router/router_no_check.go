package router

import (
	"github.com/gin-gonic/gin"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/system"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/ginx"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, noCheckProRouter)
}

func noCheckProRouter(r *gin.RouterGroup) {
	v1 := r.Group("/")
	{
		template := system.Template{}
		user := system.NewUser()

		v1.POST("/t_add", ginx.Handle(template.Add))
		v1.POST("/t_a_add", ginx.Handle(template.AbilityAdd))

		v1.POST("/login", ginx.Handle(user.Login))
	}
}
