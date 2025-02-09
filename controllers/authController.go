package controllers

import (
	"campus_secondhand/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 注册
func RegisterHandler(c *gin.Context) {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "输入无效"})
		return
	}

	err := services.RegisterUser(request.Username, request.Password)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}

// 登录
func LoginHandler(c *gin.Context) {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "输入无效"})
		return
	}

	user, err := services.AuthenticateUser(request.Username, request.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "登录成功", "userID": user.ID})
}

// 退出登录
func LogoutHandler(c *gin.Context) {
	var request struct {
		UserID uint `json:"user_id"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "输入无效"})
		return
	}

	err := services.LogoutUser(request.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "退出失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "退出成功"})
}
