package main

import (
	"campus_secondhand/controllers"
	"campus_secondhand/database"
	"campus_secondhand/models"
	"github.com/gin-gonic/gin"
)

func main() {
	//初始化数据库
	database.InitDB()
	// 执行数据库迁移，确保表被创建
	models.Migrate()
	//注册路由
	router := gin.Default()

	router.POST("/register", controllers.RegisterHandler)
	router.POST("/login", controllers.LoginHandler)

	router.Run(":8080")
}
