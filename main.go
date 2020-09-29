package main

import (
	"fmt"
	"gmt-go/conf/setting"
	models "gmt-go/model"
	"gmt-go/router"
	"strconv"
)

func init() {
	models.Setup()
}

func main() {
	setting.InitSetting()
	r := router.SetupRouter()
	fmt.Println(setting.Conf().Server.HTTPPort)
	r.Run(":" + strconv.Itoa(setting.Conf().Server.HTTPPort)) // listen and serve on 0.0.0.0:1234 (for windows "localhost:1234")
}
