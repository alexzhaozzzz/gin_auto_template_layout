package router

import (
	"github.com/gin-gonic/gin"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/ginx"
)

func init() {
	routerCheckRole = append(routerCheckRole, checkProRouter)
}

func checkProRouter(r *gin.RouterGroup, gHanFun ...gin.HandlerFunc) {
	v1 := r.Group("/").Use(gHanFun...)
	{
		//user := logic.NewUser()
		menu := logic.NewMenu()
		merchant := logic.NewMerchant()
		role := logic.NewRole()
		permission := logic.NewPermissions()

		//v1.GET("/user", ginx.Handle(user.GetList))
		//v1.POST("/user/edit", ginx.Handle(user.Edit))
		//v1.POST("/user/delete", ginx.Handle(user.Delete))

		v1.GET("/menu", ginx.Handle(menu.GetList))
		v1.POST("/menu/add", ginx.Handle(menu.Add))
		v1.POST("/menu/edit", ginx.Handle(menu.Edit))
		v1.POST("/menu/delete", ginx.Handle(menu.Delete))

		v1.GET("/role", ginx.Handle(role.GetList))
		v1.POST("/role/edit", ginx.Handle(role.Edit))
		v1.POST("/role/delete", ginx.Handle(role.Delete))

		v1.GET("/role/user", ginx.Handle(role.GetUserList))
		v1.POST("/role/user/edit", ginx.Handle(role.EditUser))
		v1.POST("/role/user/delete", ginx.Handle(role.DeleteUser))

		v1.GET("/role/menu", ginx.Handle(role.GetMenuList))
		v1.POST("/role/menu/edit", ginx.Handle(role.EditMenu))
		v1.POST("/role/menu/delete", ginx.Handle(role.DeleteMenu))

		v1.GET("/role/permission", ginx.Handle(role.GetPermissionsList))
		v1.POST("/role/permission/edit", ginx.Handle(role.EditPermissions))
		v1.POST("/role/permission/delete", ginx.Handle(role.DeletePermissions))

		v1.GET("/merchant", ginx.Handle(merchant.GetList))
		v1.GET("/permission", ginx.Handle(permission.GetList))
	}
}
