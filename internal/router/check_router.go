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
		user := logic.NewUser()
		menu := logic.NewMenu()
		merchant := logic.NewMerchant()
		role := logic.NewRole()
		permission := logic.NewPermissions()
		roleusers := logic.NewRoleUsers()
		rolemenu := logic.NewRoleMenu()
		rolepermission := logic.NewRolePermissions()

		v1.GET("/user", ginx.Handle(user.GetList))
		v1.POST("/user/edit", ginx.Handle(user.Edit))
		v1.POST("/user/delete", ginx.Handle(user.Delete))

		v1.GET("/menu", ginx.Handle(menu.GetList))
		v1.POST("/menu/add", ginx.Handle(menu.Add))
		v1.POST("/menu/edit", ginx.Handle(menu.Edit))
		v1.POST("/menu/delete", ginx.Handle(menu.Delete))

		v1.GET("/role", ginx.Handle(role.GetList))
		v1.POST("/role/add", ginx.Handle(role.Add))
		v1.POST("/role/edit", ginx.Handle(role.Edit))
		v1.POST("/role/delete", ginx.Handle(role.Delete))

		v1.GET("/roleusers", ginx.Handle(roleusers.GetList))
		v1.POST("/roleusers/add", ginx.Handle(roleusers.Add))
		v1.POST("/roleusers/edit", ginx.Handle(roleusers.Edit))
		v1.POST("/roleusers/delete", ginx.Handle(roleusers.Delete))

		v1.GET("/rolemenu", ginx.Handle(rolemenu.GetList))
		v1.POST("/rolemenu/add", ginx.Handle(rolemenu.Add))
		v1.POST("/rolemenu/edit", ginx.Handle(rolemenu.Edit))
		v1.POST("/rolemenu/delete", ginx.Handle(rolemenu.Delete))
		//
		v1.GET("/rolepermission", ginx.Handle(rolepermission.GetList))
		v1.POST("/rolepermission/add", ginx.Handle(rolepermission.Add))
		v1.POST("/rolepermission/edit", ginx.Handle(rolepermission.Edit))
		v1.POST("/rolepermission/delete", ginx.Handle(rolepermission.Delete))

		v1.GET("/merchant", ginx.Handle(merchant.GetList))
		v1.GET("/merchant/add", ginx.Handle(merchant.Add))
		v1.GET("/merchant/edit", ginx.Handle(merchant.Edit))
		v1.GET("/merchant/delete", ginx.Handle(merchant.Delete))

		v1.GET("/permission", ginx.Handle(permission.GetList))
		v1.GET("/permission/add", ginx.Handle(permission.Add))
		v1.GET("/permission/edit", ginx.Handle(permission.Edit))
		v1.GET("/permission/delete", ginx.Handle(permission.Delete))
	}
}
