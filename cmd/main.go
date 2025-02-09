package main

import (
	"campus_secondhand/database"
	"campus_secondhand/models"
	"campus_secondhand/routes"
)

func main() {
	database.InitMySQL()
	database.InitRedis()
	database.DB.AutoMigrate(&models.User{}, &models.Item{})

	r := routes.SetupRouter()
	r.Run(":8080")
}
