package router

import (
	"github.com/gin-gonic/gin"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/logic/ability"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/pkg/ginx"
)

func init() {
	routerCheckRole = append(routerCheckRole, statisticsRouter)
}

func statisticsRouter(r *gin.RouterGroup, gHanFun ...gin.HandlerFunc) {
	v1 := r.Group("/").Use(gHanFun...)
	{
		c := ability.NewChannel()
		r := ability.NewRetain()
		lvt := ability.NewLvt()
		m := ability.NewManufacturers()
		t := ability.NewTopic()
		ci := ability.NewChannelInfo()

		v1.GET("/statistics/channel", ginx.Handle(c.GetListByPage))
		v1.GET("/statistics/retain", ginx.Handle(r.GetListByPage))
		v1.GET("/statistics/lvt", ginx.Handle(lvt.GetListByPage))
		v1.GET("/statistics/manufacturers", ginx.Handle(m.GetListByPage))
		v1.GET("/statistics/topic", ginx.Handle(t.GetListByPage))
		v1.GET("/statistics/channelinfo", ginx.Handle(ci.GetListByPage))
	}
}
