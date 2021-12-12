package database

import (
	"fmt"
	"github.com/dora-exku/shorturl/pkg/config"

	gormlogger "gorm.io/gorm/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ContentDB() *gorm.DB {
	// 错误
	var err error
	// 定义数据库链接变量
	var (
		host     = config.GetString("database.mysql.host")
		port     = config.GetString("database.mysql.port")
		database = config.GetString("database.mysql.database")
		username = config.GetString("database.mysql.username")
		password = config.GetString("database.mysql.password")
		charset  = config.GetString("database.mysql.charset")
	)
	// 数据库信息
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s", username, password, host, port, database, charset, true, "Local")
	mysqlConfig := mysql.New(mysql.Config{
		DSN: dsn,
	})
	// gorm 日志等级
	var level gormlogger.LogLevel
	// 判断是否debug
	if config.GetBool("app.debug") {
		level = gormlogger.Info
	} else {
		level = gormlogger.Error
	}
	// 打开数据库链接
	DB, err = gorm.Open(mysqlConfig, &gorm.Config{
		Logger: gormlogger.Default.LogMode(level),
	})
	if err != nil {
		fmt.Println(err)
	}
	return DB
}
