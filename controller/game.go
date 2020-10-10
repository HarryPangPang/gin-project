package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 获取游戏
func GetAllGames(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "获取游戏权限成功",
		"data": nil,
	})
}
