package database

import (
	"log"

	"campus_secondhand/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitMySQL 初始化 MySQL 数据库连接
func InitMySQL() {
	var err error
	dsn := "root:root@tcp(localhost:3306)/campus_secondhand?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}
	log.Println("MySQL 连接成功")

	// 自动迁移 User 表
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("MySQL 表迁移失败: %v", err)
	}
	log.Println("MySQL 表迁移成功")
}
