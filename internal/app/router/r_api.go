package router

import (
	"github.com/gin-gonic/gin"
	"github.com/peanut-pg/gin_admin/internal/app/middleware"
)

// RegisterAPI register api group router
func (a *Router) RegisterAPI(app *gin.Engine) {
	g := app.Group("/api")
	g.Use(middleware.UserAuthMiddleware(a.Auth,
		middleware.AllowPathPrefixSkipper("/api/v1/pub/login"),
	))

	g.Use(middleware.RateLimiterMiddleware())

	v1 := g.Group("/v1")
	{
		gUser := v1.Group("users")
		gUser.GET("", a.UserAPI.Query)
		gUser.GET(":id", a.UserAPI.Get)
		gUser.POST("", a.UserAPI.Create)
		gUser.PUT(":id", a.UserAPI.Update)
		gUser.DELETE(":id", a.UserAPI.Delete)
		gUser.PATCH(":id/enable", a.UserAPI.Enable)
		gUser.PATCH(":id/disable", a.UserAPI.Disable)
	}

}
