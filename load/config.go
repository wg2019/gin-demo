// Package load 加载配置
package load

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

func init() {
	cfg, err := ini.Load("config/app.ini")
	if err != nil {
		log.Fatalf("load.config, fail to parse 'conf/app.ini': %v", err)
	}
	mapTo(cfg, "server", &ServerSetting)
	mapTo(cfg, "database", &DatabaseSetting)
	mapTo(cfg, "redis", &RedisSetting)
}

// mapTo map section
func mapTo(cfg *ini.File, section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}

// Server 服务配置
type Server struct {
	RunMode      string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// ServerSetting 辅助配置
var ServerSetting = &Server{}

// Database 数据库配置
type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

// DatabaseSetting 数据库配置
var DatabaseSetting = &Database{}

// Redis redis配置
type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

// RedisSetting redis配置
var RedisSetting = &Redis{}
