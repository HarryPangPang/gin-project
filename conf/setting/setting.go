package setting

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

//相应设置配置
type Setting struct {
	RunMode  string   `yaml:"runMode"`
	Server   server   `yaml:"server"`
	Database database `yaml:"database"`
}

//服务配置
type server struct {
	HTTPPort     int           `yaml:"HTTPPort"`
	ReadTimeout  time.Duration `yaml:"readTimeout"`
	WriteTimeout time.Duration `yaml:"writeTimeout"`
}

//数据库配置
type database struct {
	Type     string `yaml:"type"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	DBName   string `yaml:"dbname"`
}

// 微信登陆

var conf = &Setting{}

//初始化方法
func InitSetting() {

	env := os.Getenv("GO_ENV")
	fmt.Println("读取当前环境" + env)
	yamlFile, err := ioutil.ReadFile("etc/" + env + ".yaml")
	if err != nil {
		fmt.Println(err)
		fmt.Println("获取错误处理")
	}

	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		fmt.Println(err)
		fmt.Println("获取错误处理")
	}
}

//获取配置  外部调用使用
func Conf() *Setting {
	return conf
}
