package middleware

//import (
//	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/ent"
//	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/auth"
//	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/bootstrap"
//	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/ginx"
//	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/statusx"
//	"context"
//	"github.com/casbin/casbin/v2/util"
//	"net/http"
//
//	"github.com/gin-gonic/gin"
//	"github.com/go-admin-team/go-admin-core/sdk/pkg/response"
//)
//
//type UrlInfo struct {
//	Url    string
//	Method string
//}
//
//// CasbinExclude casbin 排除的路由列表
//var CasbinExclude = []UrlInfo{
//	{Url: "/api/v1/dict/type-option-select", Method: "GET"},
//}
//
//// AuthCheckRole 权限检查中间件
//func AuthCheckRole(app *server.App) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		gxCtx := ginx.NewContext(c)
//		uid := auth.GetJwtUserId(gxCtx)
//		if uid <= 0 {
//			gxCtx.Render(statusx.StatusUnauthorized, "nil", http.StatusUnauthorized)
//			c.Abort()
//			return
//		}
//
//		uinfo, err := getUser(app, uid)
//		if err != nil {
//			gxCtx.Render(statusx.StatusPermissionsError, "nil", http.StatusUnauthorized)
//			c.Abort()
//			return
//		}
//
//		checkRole := isSuperAdmin(app, 1)
//		//检查权限
//		if checkRole {
//			c.Next()
//			return
//		}
//
//		var exclude bool
//		for _, i := range CasbinExclude {
//			if util.KeyMatch2(c.Request.URL.Path, i.Url) && c.Request.Method == i.Method {
//				exclude = true
//				break
//			}
//		}
//
//		if exclude {
//			c.Next()
//			return
//		}
//
//		var res bool
//		res, err = e.Enforce(v["rolekey"], c.Request.URL.Path, c.Request.Method)
//		if err != nil {
//			response.Error(c, 500, err, "")
//			return
//		}
//
//		if res {
//			c.Next()
//		} else {
//			c.JSON(http.StatusOK, gin.H{
//				"code": 403,
//				"msg":  "对不起，您没有该接口访问权限，请联系管理员",
//			})
//			c.Abort()
//			return
//		}
//	}
//}
//
//func isSuperAdmin(app *server.App, roleId uint64) bool {
//	info, err := app.Data.DB.SysRoles.Get(context.Background(), int64(roleId))
//	if err != nil {
//		return false
//	}
//
//	isSuper := false
//	if info.Code == "SuperAdmin" || info.Code == "admin" {
//		isSuper = true
//	}
//
//	return isSuper
//}
//
//func getUser(app *server.App, uId uint64) (*ent.User, error) {
//	return app.Data.DB.User.Get(context.Background(), int64(uId))
//}
