package setting

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
)

// 定义服务器配置结构体
type Server struct {
	Ip   string
	Port string
}
type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var ServerSetting = &Server{}
var DatabaseSetting = &Database{}
var config *ini.File

func Setup() {
	var err error
	dir, _ := os.Getwd()
	config, err = ini.Load(dir + "/config/app.ini")
	if err != nil {
		log.Fatal("Fail to parse 'config/app.ini': %v", err)
	}
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
}

func mapTo(section string, v interface{}) {
	err := config.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}