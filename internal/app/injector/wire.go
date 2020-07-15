package injector

import (
	"github.com/google/wire"
	"github.com/peanut-pg/gin_admin/internal/app/router"
)

// BuildInjector 生成注入器
func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		InitGormDB,
		InitAuth,
		InitGinEngine,
		InjectorSet,
		router.RouterSet,
	)
	return new(Injector), nil, nil
}
