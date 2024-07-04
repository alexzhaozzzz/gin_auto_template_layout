package router

import (
	"github.com/gin-gonic/gin"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/ability"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/ginx"
)

func init() {
	routerCheckRole = append(routerCheckRole, userRouter)
}

func userRouter(r *gin.RouterGroup, gHanFun ...gin.HandlerFunc) {
	v1 := r.Group("/").Use(gHanFun...)
	{
		goldc := ability.NewGoldChanges()
		gameLog := ability.NewGameLog()
		login := ability.NewLoginLog()
		payouts := ability.NewPayouts()
		rechargeLog := ability.NewRechargeLog()
		toc := ability.NewThirdOrderCheck()
		uf := ability.NewUserFight()
		ub := ability.NewUserBet()

		v1.GET("/ability/goldchange", ginx.Handle(goldc.GetListByPage))
		v1.GET("/ability/gamelog", ginx.Handle(gameLog.GetListByPage))
		v1.GET("/ability/login", ginx.Handle(login.GetListByPage))
		v1.GET("/ability/payoutsaudit", ginx.Handle(payouts.GetAuditListByPage))
		v1.GET("/ability/payouts", ginx.Handle(payouts.GetListByPage))
		v1.GET("/ability/baseinfo", ginx.Handle(payouts.GetUserBaseInfo))
		v1.GET("/ability/userpayouts", ginx.Handle(payouts.GetUserListByPage))
		v1.GET("/ability/rechargelog", ginx.Handle(rechargeLog.GetListByPage))
		v1.GET("/ability/checkorderstatus", ginx.Handle(toc.CheckOrderStatus))
		v1.GET("/ability/userfight", ginx.Handle(uf.GetListByPage))
		v1.GET("/ability/userbet", ginx.Handle(ub.GetListByPage))
	}
}
