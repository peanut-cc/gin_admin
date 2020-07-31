package injector

import (
	"context"
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist"
	"github.com/peanut-pg/gin_admin/internal/app/config"
	"github.com/peanut-pg/gin_admin/pkg/logger"
)

// InitCasbin 初始化casbin
func InitCasbin(adapter persist.Adapter) (*casbin.SyncedEnforcer, func(), error) {
	cfg := config.C.Casbin
	if cfg.Model == "" {
		return new(casbin.SyncedEnforcer), nil, nil
	}

	e, err := casbin.NewSyncedEnforcer(cfg.Model)
	if err != nil {
		return nil, nil, err
	}
	e.EnableLog(cfg.Debug)
	logger.Infof(context.Background(), "init InitWithModelAndAdapter")
	err = e.InitWithModelAndAdapter(e.GetModel(), adapter)
	logger.Infof(context.Background(), "stop InitWithModelAndAdapter")
	if err != nil {
		return nil, nil, err
	}
	e.EnableEnforce(cfg.Enable)

	cleanFunc := func() {}
	if cfg.AutoLoad {
		e.StartAutoLoadPolicy(time.Duration(cfg.AutoLoadInternal) * time.Second)
		cleanFunc = func() {
			e.StopAutoLoadPolicy()
		}
	}

	return e, cleanFunc, nil
}
