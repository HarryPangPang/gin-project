package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// SetLang 设置语言信息
func SetLang(c *gin.Context) {
	lang := c.DefaultQuery("lang", "en")
	session := sessions.Default(c)
	session.Set("lang", lang)
	err := session.Save()
	if err != nil {
		logger.Errorln("session保存失败", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 1,
			"msg":  "设置语言失败",
			"data": string(err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "设置语言成功",
		"data": nil,
	})
}

//GetLang 获取当前语言
func GetLang(c *gin.Context) {
	session := sessions.Default(c)
	lang := session.Get("lang")
	log.Println("lang", lang)
	if lang == nil {
		lang = "en"
		fmt.Println(lang)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "语言成功",
		"data": lang,
	})
}
