package dao

import (
	"fmt"
	gorm "gorm.io/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

type Dao struct {
}

var db *gorm.DB

func Db() *gorm.DB {
	if db == nil {
		db = initDb()
	}
	return db
}

func init() {
	fmt.Printf("数据库初始化....\n")
	if db == nil {
		db = initDb()
	}
}
func initDb() *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:              time.Second,   // Slow SQL threshold
			LogLevel:                   logger.Info, // Log level
			IgnoreRecordNotFoundError: false,           // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,          // Disable color
		},
	)

	fmt.Printf("数据库连接.....\n")
	dsn := "golang:123456@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		fmt.Printf("输出库连接异常 %s \n",err.Error())
	}

	sqlDB, _ := db.DB()

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}
