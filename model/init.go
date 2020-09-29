package model

import (
	"fmt"
	"gmt-go/conf/setting"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Setup() {
	user := setting.Conf().Database.User
	host := setting.Conf().Database.Host
	password := setting.Conf().Database.Password
	dbName := setting.Conf().Database.DBName
	var err error
	dsn := user + ":" + password + "@tcp(" + host + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if DB.Error != nil {
		fmt.Printf("database error %v", DB.Error)
	}
	AutoMigrate()
}

// 自动迁移
func AutoMigrate() {
	DB.AutoMigrate(&User{})
}
