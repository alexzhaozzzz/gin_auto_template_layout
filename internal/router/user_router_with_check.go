package router

import (
	"github.com/gin-gonic/gin"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/billing"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/user"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/ginx"
)

func init() {
	routerCheckRole = append(routerCheckRole, userRouter)
}

func userRouter(r *gin.RouterGroup, gHanFun ...gin.HandlerFunc) {
	v1 := r.Group("/").Use(gHanFun...)
	{
		goldc := user.NewGoldChanges()
		gameLog := user.NewGameLog()
		login := user.NewLogin()
		payouts := billing.NewPayouts()
		rechargeLog := billing.NewRechargeLog()

		v1.GET("/ability/goldchange", ginx.Handle(goldc.GetListByPage))
		v1.GET("/ability/gamelog", ginx.Handle(gameLog.GetListByPage))
		v1.GET("/ability/login", ginx.Handle(login.GetListByPage))
		v1.GET("/ability/payoutsaudit", ginx.Handle(payouts.GetAuditListByPage))
		v1.GET("/ability/payouts", ginx.Handle(payouts.GetListByPage))
		v1.GET("/ability/rechargelog", ginx.Handle(rechargeLog.GetListByPage))
	}
}
