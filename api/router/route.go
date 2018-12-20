package router

import (
	"app/handler"

	"app/router/middleware"
	"github.com/gin-gonic/gin"
	"github.com/ilibs/very/core"
	"github.com/ilibs/very/debug"
)

func Route(router *gin.Engine) {
	//设置模板
	//core.SetTemplate(router)

	//中间件
	router.Use(core.Middleware(middleware.Ginrus))
	router.HEAD("/heartbeat/check", core.Handle(handler.Check))

	router.POST("/api/login", core.Handle(handler.Login))

	admin := router.Group("/api/admin")
	admin.Use(core.Middleware(middleware.AuthLogin))
	{
		admin.POST("/user", core.Handle(handler.AdminUser))
		admin.POST("/list", core.Handle(handler.AdminMoodList))
		admin.POST("/add", core.Handle(handler.AdminMoodPost))
		admin.POST("/delete", core.Handle(handler.AdminMoodDelete))
	}

	//debug handler
	debug.Route(router)
}
