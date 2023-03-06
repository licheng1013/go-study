package common

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
)

var App = gin.New()

// Db 全局数据库操作实例
var Db *gorm.DB

// AppConfig 全局配置变量
var AppConfig *SystemConfig

type SystemConfig struct {
	MongoUrl    string      `yaml:"mongoUrl"`
	MysqlUrl    string      `yaml:"mysqlUrl"`
	RedisConfig RedisConfig `yaml:"redisConfig"`
	FilePath    string      `yaml:"filePath"`
	ExcludePath []string    `yaml:"excludePath"`
	Addr        string      `yaml:"addr"`
}

type RedisConfig struct {
	User string   `yaml:"user"`
	Pass string   `yaml:"pass"`
	List []string `yaml:"list"`
}

func ParseAppConfig(data []byte) {
	if len(data) == 0 {
		panic("配置为空!")
	}
	_ = yaml.Unmarshal(data, &AppConfig)
}
