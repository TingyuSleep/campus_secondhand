package routes

import (
	"campus_secondhand/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/register", controllers.RegisterHandler)
	r.POST("/login", controllers.LoginHandler)
	r.POST("/logout", controllers.LogoutHandler)
	return r
}
