package middleware

import (
	"gmt-go/conf/setting"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// KEY gin session key
const KEY = "gmt-go-secrert"

//EnableCookieSession 使用 Cookie 保存 session
func EnableCookieSession() gin.HandlerFunc {
	store := cookie.NewStore([]byte(KEY))
	return sessions.Sessions("go-gmt-session", store)
}

//AuthSessionMiddle session验证中间件
func AuthSessionMiddle() gin.HandlerFunc {
	now := time.Now().Unix()
	serverURL := setting.Conf().WeixinOauth.ServerUrl
	accessKey := setting.Conf().WeixinOauth.AccessKey
	redirectURL := setting.Conf().WeixinOauth.RedirectURL
	oauthRedirectURL := serverURL + "/user/login" + "?accessKey=" + accessKey + "&redirectURL=" + redirectURL
	return func(c *gin.Context) {
		session := sessions.Default(c)
		accessTokenExpire := session.Get("AccessTokenExpire")
		path := c.Request.URL.Path
		log.Println("accessTokenExpire", accessTokenExpire)
		if accessTokenExpire != nil {
			accessTokenExpiresInt := accessTokenExpire.(int64)
			if accessTokenExpiresInt < now {
				session.Clear()
				session.Save()
				c.Redirect(http.StatusMovedPermanently, oauthRedirectURL)
				return
			}
		}
		if accessTokenExpire == nil {
			if path == "/api/auth/login" {
				c.Next()
				return
			}
			session.Clear()
			session.Save()
			c.Redirect(http.StatusMovedPermanently, oauthRedirectURL)
			return
		}
		c.Next()
		return
	}
}

//SaveAuthSession 注册和登陆时都需要保存session信息
func SaveAuthSession(c *gin.Context, id uint) {
	session := sessions.Default(c)
	session.Set("userId", id)
	session.Save()
}

//ClearAuthSession 退出时清除session
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
		log.Println(userId)
		// userName = models.UserDetail(userId).Name
	}
	data := make(map[string]interface{})
	data["hasSession"] = hasSession
	data["userName"] = userName
	return data
}
