package router

import (
	"app/handler"

	"app/router/middleware"
	"github.com/gin-gonic/gin"
	"github.com/goapt/gee"
)

func Route(router *gin.Engine) {
	//设置模板
	//core.SetTemplate(router)

	//中间件
	router.Use(gee.Middleware(middleware.Ginrus))
	router.HEAD("/heartbeat/check", gee.Handle(handler.Check))

	router.POST("/api/login", gee.Handle(handler.Login))

	admin := router.Group("/api/admin")
	admin.Use(gee.Middleware(middleware.AuthLogin))
	{
		admin.POST("/user", gee.Handle(handler.AdminUser))
		admin.POST("/list", gee.Handle(handler.AdminMoodList))
		admin.POST("/add", gee.Handle(handler.AdminMoodPost))
		admin.POST("/delete", gee.Handle(handler.AdminMoodDelete))
	}

	//debug handler
	gee.DebugRoute(router)
}
