package controller

import (
	"encoding/json"
	"fmt"
	"gmt-go/helper"
	servicegame "gmt-go/service/game"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	logger *logrus.Logger
)

func init() {
	logger = helper.Logger()
}

// GetAllGames 获取游戏
func GetAllGames(c *gin.Context) {
	session := sessions.Default(c)
	sessionGame := session.Get("game")
	if sessionGame != nil {
		strSessionGame := fmt.Sprintf("%v", sessionGame)
		res := make(map[string]interface{}, 0)
		if err := json.Unmarshal([]byte(strSessionGame), &res); err != nil {
			logger.Errorln(err)
			c.JSON(http.StatusBadGateway, gin.H{
				"code": 1,
				"msg":  "获取游戏session失败",
				"data": string(err.Error()),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "获取游戏session成功",
			"data": res,
		})
		return
	}
	results, err := servicegame.QueryAllGames()
	if err != nil {
		logger.Errorln(err)
		c.JSON(http.StatusBadGateway, gin.H{
			"code": 1,
			"msg":  "获取游戏失败",
			"data": string(err.Error()),
		})
		return
	}

	res := make(map[string]interface{})
	cacheSlice1 := make([]interface{}, 0)
	cacheSlice2 := make([]interface{}, 0)
	for i := 0; i < len(results); i++ {
		result := results[i]
		cacheMap1 := make(map[string]interface{}, 0)
		cacheMap2 := make(map[string]interface{}, 0)
		for k1, v1 := range result {
			if strings.Contains(k1, ".") {
				k1TableColumn := strings.Split(k1, ".")
				k1Table := k1TableColumn[0]
				k1Column := k1TableColumn[1]
				if res[k1Table] == nil {
					res[k1Table] = make([]interface{}, 0)
				}

				if k1Table == "channelAppIds" && v1 != "" {
					cacheMap1[k1Column] = v1
				}
				if k1Table == "gameAppIds" && v1 != "" {
					cacheMap2[k1Column] = v1
				}
			} else {
				res[k1] = v1
			}
		}
		if len(cacheMap1) != 0 {
			cacheSlice1 = append(cacheSlice1, cacheMap1)
		}
		if len(cacheMap2) != 0 {
			cacheSlice2 = append(cacheSlice2, cacheMap2)
		}

		res["channelAppIds"] = cacheSlice1
		res["gameAppIds"] = cacheSlice2
	}
	resString, _ := json.Marshal(res)
	session.Set("game", string(resString))
	errSession := session.Save()
	if errSession != nil {
		fmt.Println(errSession)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "获取游戏成功",
		"data": res,
	})
}
