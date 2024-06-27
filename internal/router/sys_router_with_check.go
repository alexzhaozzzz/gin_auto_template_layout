package router

import (
	"github.com/gin-gonic/gin"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/system"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/ginx"
)

func init() {
	routerCheckRole = append(routerCheckRole, checkProRouter)
}

func checkProRouter(r *gin.RouterGroup, gHanFun ...gin.HandlerFunc) {
	v1 := r.Group("/").Use(gHanFun...)
	{
		user := system.NewUser()
		menu := system.NewMenu()
		merchant := system.NewMerchant()
		role := system.NewRole()
		permission := system.NewPermissions()
		roleusers := system.NewRoleUsers()
		rolemenu := system.NewRoleMenu()
		rolepermission := system.NewRolePermissions()

		v1.GET("/user", ginx.Handle(user.GetListByPage))
		v1.GET("/user/info/:id", ginx.Handle(user.GetInfo))
		v1.POST("/user/add", ginx.Handle(user.Add))
		v1.POST("/user/edit", ginx.Handle(user.Edit))
		v1.POST("/user/delete", ginx.Handle(user.Delete))

		v1.GET("/menu", ginx.Handle(menu.GetList))
		v1.GET("/menu/info/:id", ginx.Handle(menu.GetInfo))
		v1.POST("/menu/add", ginx.Handle(menu.Add))
		v1.POST("/menu/edit", ginx.Handle(menu.Edit))
		v1.POST("/menu/delete", ginx.Handle(menu.Delete))

		v1.GET("/role", ginx.Handle(role.GetListByPage))
		v1.GET("/role/all", ginx.Handle(role.GetList))
		v1.GET("/role/info/:id", ginx.Handle(role.GetInfo))
		v1.POST("/role/add", ginx.Handle(role.Add))
		v1.POST("/role/edit", ginx.Handle(role.Edit))
		v1.POST("/role/delete", ginx.Handle(role.Delete))

		v1.GET("/merchant", ginx.Handle(merchant.GetListByPage))
		v1.GET("/merchant/info/:id", ginx.Handle(merchant.GetInfo))
		v1.POST("/merchant/add", ginx.Handle(merchant.Add))
		v1.POST("/merchant/edit", ginx.Handle(merchant.Edit))
		v1.POST("/merchant/delete", ginx.Handle(merchant.Delete))

		v1.GET("/permission", ginx.Handle(permission.GetList))
		v1.GET("/roleusers/:role_id", ginx.Handle(roleusers.GetList))
		v1.GET("/rolemenu/:role_id", ginx.Handle(rolemenu.GetList))
		v1.GET("/rolepermission/:role_id", ginx.Handle(rolepermission.GetList))
	}
}
