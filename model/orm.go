package model

import (
	"gmt-go/conf/setting"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"xorm.io/core"
)

var (
	// DBM 主数据库
	DBM *xorm.Engine
	// DBS 从数据库
	DBS *xorm.Engine
)

// Init 初始化数据库
func Init() {
	var err error
	user := setting.Conf().Database.User
	host := setting.Conf().Database.Host
	password := setting.Conf().Database.Password
	dbName := setting.Conf().Database.DBName
	dsnm := user + ":" + password + "@tcp(" + host + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Asia%2fShanghai"
	DBM, err = xorm.NewEngine("mysql", dsnm)
	DBM.SetMaxIdleConns(10)
	DBM.SetMaxOpenConns(200)
	// DBM.ShowSQL(true)
	// DBM.ShowExecTime(true)
	DBM.SetColumnMapper(core.SameMapper{})
	if err != nil {
		log.Printf("连接主数据库错误: %v", err)
		os.Exit(1)
	}
	//从库添加
	users := setting.Conf().DatabaseSlave.User
	hosts := setting.Conf().DatabaseSlave.Host
	passwords := setting.Conf().DatabaseSlave.Password
	dbNames := setting.Conf().DatabaseSlave.DBName
	dsns := users + ":" + passwords + "@tcp(" + hosts + ")/" + dbNames + "?charset=utf8&parseTime=True&loc=Asia%2fShanghai"
	DBS, err := xorm.NewEngine("mysql", dsns)
	// DBS.SetMaxIdleConns(10)
	// DBS.SetMaxOpenConns(200)
	DBS.ShowSQL(true)
	DBM.SetColumnMapper(core.SameMapper{})
	DBS.ShowExecTime(true)
	if err != nil {
		log.Printf("连接从数据库错误: %v", err)
		os.Exit(1)
	}
	log.Println(dsnm)
	log.Println("数据库初始化完成")

}

// GetMaster 主数据库
func GetMaster() *xorm.Engine {
	return DBM
}

// GetSlave 从数据库
func GetSlave() *xorm.Engine {
	return DBS
}
