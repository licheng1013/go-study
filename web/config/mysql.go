package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"web/common"
)

// 初始化mysql链接
func init() {
	//mysqlInit()
}

func MysqlInit() {
	// 日志打印
	newLogger := logger.Default

	log.Println("Mysql:初始化！")
	dsn := common.AppConfig.MysqlUrl
	v, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_", // 定义表前缀
			SingularTable: true, // true不在表后面+ s，
		},
	})
	if err != nil {
		log.Panic(err)
	}
	common.Db = v
}
