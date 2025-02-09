package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitMySQL() {
	//dsn,data source name,数据源名称
	dsn := "root:root@tcp(localhost:3306)/campus_secondhand?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Mysql连接失败：%v", err)
	}
	fmt.Println("Mysql连接成功")
}
