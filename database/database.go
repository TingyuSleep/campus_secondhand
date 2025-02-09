package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB() {
	dsn := "root:root@tcp(127.0.0.1:3306)/campus_secondhand?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("获取数据库连接失败:", err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatal("数据库Ping失败:", err)
	}

	fmt.Println("数据库连接成功")
}
