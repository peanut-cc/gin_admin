//+build wireinject
//The build tag makes sure the stub is not built in the final build.

package injector

import (
	"github.com/google/wire"
	"github.com/peanut-pg/gin_admin/internal/app/api"
	"github.com/peanut-pg/gin_admin/internal/app/bll/impl/bll"
	gormModel "github.com/peanut-pg/gin_admin/internal/app/model/impl/gorm/model"
	"github.com/peanut-pg/gin_admin/internal/app/module/adapter"
	"github.com/peanut-pg/gin_admin/internal/app/router"
)

// BuildInjector 生成注入器
func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		InitGormDB,
		InitAuth,
		InitGinEngine,
		InitCasbin,
		InjectorSet,
		router.RouterSet,
		gormModel.ModelSet,
		adapter.CasbinAdapterSet,
		api.APISet,
		bll.BllSet,
	)
	return new(Injector), nil, nil
}
