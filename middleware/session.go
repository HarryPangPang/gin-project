package middleware

import (
	"fmt"
	"gmt-go/conf/setting"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// gin session key
const KEY = "gmt-go-secrert"

// 使用 Cookie 保存 session
func EnableCookieSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte(KEY))
	return sessions.Sessions("gosession", store)
}

// session验证中间件
func AuthSessionMiddle() gin.HandlerFunc {
	now := time.Now().Unix()
	serverUrl := setting.Conf().WeixinOauth.ServerUrl
	accessKey := setting.Conf().WeixinOauth.AccessKey
	redirectURL := setting.Conf().WeixinOauth.RedirectURL
	oauthRedirectURL := serverUrl + "/user/login" + "?accessKey=" + accessKey + "&redirectURL=" + redirectURL
	return func(c *gin.Context) {
		session := sessions.Default(c)
		accessTokenExpire := session.Get("AccessTokenExpire")
		accessTokenExpiresInt := accessTokenExpire.(int64)
		if accessTokenExpire == nil || accessTokenExpiresInt < now {
			c.Redirect(http.StatusMovedPermanently, oauthRedirectURL)
			return
		}
		c.Next()
		return
	}
}

// 注册和登陆时都需要保存seesion信息
func SaveAuthSession(c *gin.Context, id uint) {
	session := sessions.Default(c)
	session.Set("userId", id)
	session.Save()
}

// 退出时清除session
func ClearAuthSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

func HasSession(c *gin.Context) bool {
	session := sessions.Default(c)
	if sessionValue := session.Get("userId"); sessionValue == nil {
		return false
	}
	return true
}

func GetSessionUserId(c *gin.Context) uint {
	session := sessions.Default(c)
	sessionValue := session.Get("userId")
	if sessionValue == nil {
		return 0
	}
	return sessionValue.(uint)
}

func GetUserSession(c *gin.Context) map[string]interface{} {

	hasSession := HasSession(c)
	userName := ""
	if hasSession {
		userId := GetSessionUserId(c)
		fmt.Println(userId)
		// userName = models.UserDetail(userId).Name
	}
	data := make(map[string]interface{})
	data["hasSession"] = hasSession
	data["userName"] = userName
	return data
}
