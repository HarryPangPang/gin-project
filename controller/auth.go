package controller

import (
	"encoding/json"
	"gmt-go/conf/setting"
	"gmt-go/helper"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type code2tokenForm struct {
	appID     string
	timestamp int
	code      string
}

// TokenData 返回的所有token数据
type TokenData struct {
	AccessToken        string `json:"accessToken"`
	AccessTokenExpire  int64  `json:"accessTokenExpire"`
	RefreshToken       string `json:"refreshToken"`
	RefreshTokenExpire int64  `json:"refreshTokenExpire"`
}

// Code2tokenRes 返回结果
type Code2tokenRes struct {
	Result int       `json:"result"`
	Msg    string    `json:"msg"`
	Data   TokenData `json:"data"`
}

// TokenType token类型枚举
type TokenType int

const (
	code         TokenType = 0 //普通登陆
	refreshToken TokenType = 1 //更新token
	accessToken  TokenType = 2 //获取用户名，权限
)

// Login 登陆
func Login(c *gin.Context) {
	code := c.Query("code")
	respBody := code2token(code)
	session := sessions.Default(c)
	session.Set("AccessToken", respBody.AccessToken)
	session.Set("AccessTokenExpire", respBody.AccessTokenExpire)
	session.Set("RefreshToken", respBody.RefreshToken)
	session.Set("RefreshTokenExpire", respBody.RefreshTokenExpire)

	userinfo := token2userinfo(respBody.AccessToken)
	session.Set("wxId", userinfo["wxId"])
	session.Set("name", userinfo["name"])
	session.Set("mobile", userinfo["mobile"])
	session.Set("email", userinfo["email"])
	session.Set("avatar", userinfo["avatar"])
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "登陆成功",
		"data": nil,
	})
}

//Logout 登出
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "退出成功",
		"data": nil,
	})
}

//UserPrivs 用户权限
func UserPrivs(c *gin.Context) {
	session := sessions.Default(c)
	accessToken := session.Get("AccessToken").(string)
	userPrivs := GetUserPrivs(c, accessToken)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "获取用户权限成功",
		"data": userPrivs,
	})
}

//GetUserInfo 用户信息
func GetUserInfo(c *gin.Context) {
	session := sessions.Default(c)
	userinfo := make(map[string]string)
	userinfo["wxId"] = session.Get("wxId").(string)
	userinfo["name"] = session.Get("name").(string)
	userinfo["mobile"] = session.Get("mobile").(string)
	userinfo["email"] = session.Get("email").(string)
	userinfo["avatar"] = session.Get("avatar").(string)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "获取用户信息成功",
		"data": userinfo,
	})
}

func getSignedQuery(token string, tokenType TokenType) string {
	secretKey := setting.Conf().WeixinOauth.SecretKey
	appID := setting.Conf().WeixinOauth.AccessKey
	timestamp := time.Now().Unix()
	paramStr := "appId=" + appID + "&code=" + token + "&timestamp=" + strconv.FormatInt(timestamp, 10)
	switch tokenType {
	case 1:
		paramStr = "accessToken=" + token + "&appId=" + appID + "&timestamp=" + strconv.FormatInt(timestamp, 10)
	case 2:
		paramStr = "accessToken=" + token + "&appId=" + appID + "&timestamp=" + strconv.FormatInt(timestamp, 10)
	}
	queryString := "?" + paramStr + "&signature=" + helper.Md5Decode(paramStr+secretKey)
	return queryString
}

// token换取用户信息
func token2userinfo(accessToken string) map[string]interface{} {
	serverURL := setting.Conf().WeixinOauth.ServerUrl
	url := serverURL + "/api/oauth/user/info" + getSignedQuery(accessToken, 2)
	resp, err := http.Get(url)
	if err != nil {
		errors.Wrap(err, "请求/api/oauth/user/info失败")
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errors.Wrap(err, "token2userinfo读取失败")
	}
	res := helper.String2Map(string(content))
	data := res["data"].(map[string]interface{})
	return data
}

// 获取token
func code2token(token string) TokenData {
	serverURL := setting.Conf().WeixinOauth.ServerUrl
	now := time.Now().Unix()
	url := serverURL + "/api/oauth/code2token" + getSignedQuery(token, 0)
	resp, err := http.Get(url)
	if err != nil {
		errors.Wrap(err, "请求/oauth/code2token失败")
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errors.Wrap(err, "code2token读取失败")
	}
	var code2tokenRes Code2tokenRes
	json.Unmarshal([]byte(string(content)), &code2tokenRes)
	if code2tokenRes.Data.AccessTokenExpire < now && code2tokenRes.Data.RefreshTokenExpire > now {
		code2tokenRes = refreshCode2token(code2tokenRes.Data.RefreshToken)
	}
	return code2tokenRes.Data
}

// 更新token
func refreshCode2token(refreshToken string) Code2tokenRes {
	serverURL := setting.Conf().WeixinOauth.ServerUrl
	url := serverURL + "/api/oauth/refreshToken" + getSignedQuery(refreshToken, 1)
	resp, err := http.Get(url)
	if err != nil {
		errors.Wrap(err, "请求/oauth/refreshToken失败")
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errors.Wrap(err, "code2token读取失败")
	}
	var code2tokenRes Code2tokenRes
	json.Unmarshal([]byte(string(content)), &code2tokenRes)
	return code2tokenRes
}

//GetUserPrivs 获取菜单权限
func GetUserPrivs(c *gin.Context, accessToken string) []interface{} {
	serverURL := setting.Conf().WeixinOauth.ServerUrl
	accessKey := setting.Conf().WeixinOauth.AccessKey
	redirectURL := setting.Conf().WeixinOauth.RedirectURL
	oauthRedirectURL := serverURL + "/user/login" + "?accessKey=" + accessKey + "&redirectURL=" + redirectURL
	url := serverURL + "/api/oauth/user/privs" + getSignedQuery(accessToken, 2)
	resp, err := http.Get(url)
	if err != nil {
		errors.Wrap(err, "请求/oauth/user/privs失败")
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errors.Wrap(err, "getUserPrivs读取失败")
	}
	res := helper.String2Map(string(content))
	// FIXME: 这里有问题，会出现未登录也不跳转登陆的情况
	var data []interface{}
	if res["data"] != nil {
		data = res["data"].([]interface{})
	} else {
		log.Println(oauthRedirectURL)
	}

	return data
}

//GetAllAvaliableGames 获取游戏菜单权限
func GetAllAvaliableGames(c *gin.Context) {
	session := sessions.Default(c)
	accessToken := session.Get("AccessToken").(string)
	userPrivs := GetUserPrivs(c, accessToken)

	acc := helper.GetGameList(userPrivs)

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "获取游戏权限成功",
		"data": &acc,
	})
}
