package router

import (
	"github.com/gin-gonic/gin"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/statistics"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/ginx"
)

func init() {
	routerCheckRole = append(routerCheckRole, statisticsRouter)
}

func statisticsRouter(r *gin.RouterGroup, gHanFun ...gin.HandlerFunc) {
	v1 := r.Group("/").Use(gHanFun...)
	{
		c := statistics.NewChannel()
		lr := statistics.NewLoginRetain()
		lvt := statistics.NewLvt()
		m := statistics.NewManufacturers()
		pr := statistics.NewPayRetain()
		t := statistics.NewTopic()

		v1.GET("/statistics/channel", ginx.Handle(c.GetListByPage))
		v1.GET("/statistics/loginretain", ginx.Handle(lr.GetListByPage))
		v1.GET("/statistics/lvt", ginx.Handle(lvt.GetListByPage))
		v1.GET("/statistics/manufacturers", ginx.Handle(m.GetListByPage))
		v1.GET("/statistics/payretain", ginx.Handle(pr.GetListByPage))
		v1.GET("/statistics/topic", ginx.Handle(t.GetListByPage))
	}
}
