package auth

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gitlab.top.slotssprite.com/br_h5slots/server/merchant/internal/data/conn"
	"sync"
)

type CasbinService struct {
}

var CasbinServiceApp = new(CasbinService)

// 持久化到数据库
var (
	syncedEnforcer *casbin.SyncedEnforcer
	once           sync.Once
)

func (c *CasbinService) Casbin() *casbin.SyncedEnforcer {
	gdb := conn.GetMerchantDB()

	once.Do(func() {
		a, _ := gormadapter.NewAdapterByDB(gdb)
		syncedEnforcer, _ = casbin.NewSyncedEnforcer("./configs/rbac_model.conf", a)
	})
	_ = syncedEnforcer.LoadPolicy()
	return syncedEnforcer
}
