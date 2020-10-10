package router

import (
	"gmt-go/controller"
	"gmt-go/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	g := gin.New()
	g.Use(middleware.LoggerToFile())
	g.GET("/ping", controller.Test)

	api := g.Group("/api", middleware.EnableCookieSession())
	{

		auth := api.Group("/auth", middleware.AuthSessionMiddle())
		{
			auth.GET("/privs", controller.UserPrivs)
			auth.GET("/userinfo", controller.GetUserInfo)
			auth.GET("/login", controller.Login)
			auth.GET("/logout", controller.Logout)
			auth.GET("/menu", controller.GetAllAvaliableGames)
			auth.GET("/setlang", controller.SetLang)
			auth.GET("/getlang", controller.GetLang)
		}

		game := api.Group("/game", middleware.AuthSessionMiddle())
		{
			game.GET("/game", controller.GetAllGames)
		}
	}
	log.Println("路由初始化完成")
	return g
}
