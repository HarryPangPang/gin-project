package model

import (
	"fmt"
	"gmt-go/conf/setting"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	dbm         *xorm.Engine
	dbs         *xorm.Engine
	check_count int
)

func init() {
	var err error
	user := setting.Conf().Database.User
	host := setting.Conf().Database.Host
	password := setting.Conf().Database.Password
	dbName := setting.Conf().Database.DBName
	dsnm := user + ":" + password + "@tcp(" + host + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Asia%2fShanghai"
	dbm, err = xorm.NewEngine("mysql", dsnm)
	dbm.SetMaxIdleConns(10)
	dbm.SetMaxOpenConns(200)
	dbm.ShowSQL(true)
	dbm.ShowExecTime(true)
	if err != nil {
		fmt.Printf("Fail to connect to master: %v", err)
		os.Exit(1)
	}
	//从库添加
	users := setting.Conf().DatabaseSlave.User
	hosts := setting.Conf().DatabaseSlave.Host
	passwords := setting.Conf().DatabaseSlave.Password
	dbNames := setting.Conf().DatabaseSlave.DBName
	dsns := users + ":" + passwords + "@tcp(" + hosts + ")/" + dbNames + "?charset=utf8&parseTime=True&loc=Asia%2fShanghai"
	dbs, err := xorm.NewEngine("mysql", dsns)
	dbs.SetMaxIdleConns(10)
	dbs.SetMaxOpenConns(200)
	dbs.ShowSQL(true)
	dbs.ShowExecTime(true)
	if err != nil {
		fmt.Printf("Fail to connect to slave: %v", err)
		os.Exit(1)
	}
	fmt.Println("数据库初始化完成")
}

func GetMaster() *xorm.Engine {
	return dbm
}

func GetSlave() *xorm.Engine {
	return dbs
}
