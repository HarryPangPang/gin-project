package controller

import (
	"encoding/json"
	"gmt-go/conf/setting"
	"gmt-go/helper"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type code2tokenForm struct {
	appId     string
	timestamp int
	code      string
}

type TokenData struct {
	AccessToken        string `json: "accessToken"`
	AccessTokenExpire  int64  `json: "accessTokenExpire"`
	RefreshToken       string `json: "refreshToken"`
	RefreshTokenExpire int64  `json: "refreshTokenExpire"`
}
type Code2tokenRes struct {
	Result int       `json: "result"`
	Msg    string    `json: "msg"`
	Data   TokenData `json: "data"`
}

type TokenType int

const (
	code         TokenType = 0 //普通登陆
	refreshToken TokenType = 1 //更新token
	accessToken  TokenType = 2 //获取用户名，权限
)

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

func getSignedQuery(token string, tokenType TokenType) string {
	secretKey := setting.Conf().WeixinOauth.SecretKey
	appId := setting.Conf().WeixinOauth.AccessKey
	timestamp := time.Now().Unix()
	paramStr := "appId=" + appId + "&code=" + token + "&timestamp=" + strconv.FormatInt(timestamp, 10)
	switch tokenType {
	case 1:
		paramStr = "accessToken=" + token + "&appId=" + appId + "&timestamp=" + strconv.FormatInt(timestamp, 10)
	case 2:
		paramStr = "accessToken=" + token + "&appId=" + appId + "&timestamp=" + strconv.FormatInt(timestamp, 10)
	}
	queryString := "?" + paramStr + "&signature=" + helper.Md5Decode(paramStr+secretKey)
	return queryString
}

// token换取用户信息
func token2userinfo(accessToken string) map[string]interface{} {
	serverUrl := setting.Conf().WeixinOauth.ServerUrl
	url := serverUrl + "/api/oauth/user/info" + getSignedQuery(accessToken, 2)
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
	serverUrl := setting.Conf().WeixinOauth.ServerUrl
	now := time.Now().Unix()
	url := serverUrl + "/api/oauth/code2token" + getSignedQuery(token, 0)
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
	serverUrl := setting.Conf().WeixinOauth.ServerUrl
	url := serverUrl + "/api/oauth/refreshToken" + getSignedQuery(refreshToken, 1)
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
