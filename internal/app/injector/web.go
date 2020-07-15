package injector

import (
	"github.com/gin-gonic/gin"
	"github.com/peanut-pg/gin_admin/internal/app/config"
	"github.com/peanut-pg/gin_admin/internal/app/router"
)

// InitGinEngine 初始化gin引擎
func InitGinEngine(r router.IRouter) *gin.Engine {
	gin.SetMode(config.C.RunMode)
	app := gin.New()
	//prefixes := r.Prefixes()
	// Router register
	r.Register(app)
	return app
}
