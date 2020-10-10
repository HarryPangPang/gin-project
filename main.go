package main

import (
	"gmt-go/conf/setting"
	"gmt-go/model"
	"gmt-go/router"
	"strconv"
)

func init() {
	setting.InitSetting()
	model.Init()
	// model.Setup()
}

func main() {
	r := router.SetupRouter()
	r.Run(":" + strconv.Itoa(setting.Conf().Server.HTTPPort)) // listen and serve on 0.0.0.0:1234 (for windows "localhost:1234")
}
