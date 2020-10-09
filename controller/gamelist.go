package controller

import (
	"gmt-go/helper"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// 获取游戏菜单权限
func GetAllAvaliableGames(c *gin.Context) {
	session := sessions.Default(c)
	accessToken := session.Get("AccessToken").(string)
	userPrivs := GetUserPrivs(accessToken)
	acc := helper.GetGameList(userPrivs)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "获取游戏权限成功",
		"data": &acc,
	})
}
