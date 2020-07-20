package injector

import (
	"github.com/gin-gonic/gin"
	"github.com/peanut-pg/gin_admin/internal/app/config"
	"github.com/peanut-pg/gin_admin/internal/app/middleware"
	"github.com/peanut-pg/gin_admin/internal/app/router"
)

// InitGinEngine 初始化gin引擎
func InitGinEngine(r router.IRouter) *gin.Engine {
	gin.SetMode(config.C.RunMode)
	app := gin.New()
	app.NoMethod(middleware.NoMethodHandler())
	app.NoRoute(middleware.NoRouteHandler())

	prefixes := r.Prefixes()

	// Trace ID
	app.Use(middleware.TraceMiddleware(middleware.AllowPathPrefixNoSkipper(prefixes...)))

	// Access logger
	app.Use(middleware.LoggerMiddleware(middleware.AllowPathPrefixNoSkipper(prefixes...)))

	// Recover
	app.Use(middleware.RecoveryMiddleware())

	// Router register
	r.Register(app)
	return app
}
