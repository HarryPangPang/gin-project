package router

import (
	"gmt-go/controller"
	"gmt-go/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	g := gin.New()
	g.Use(middleware.LoggerToFile())
	g.GET("/ping", controller.Test)

	api := g.Group("/api", middleware.EnableCookieSession())
	{
		user := api.Group("/user")
		{
			user.POST("/insert", controller.AddUser)
			user.GET("/query", controller.GetUser)
		}

		auth := api.Group("/auth", middleware.AuthSessionMiddle())
		{
			auth.GET("/whoami")
			auth.GET("/login", controller.Login)
		}
	}
	return g
}
