package services

import (
	"campus_secondhand/database"
	"campus_secondhand/models"
	"errors"
)

// 注册用户
func RegisterUser(username, password string) error {
	var count int64
	database.DB.Model(&models.User{}).Where("username = ?", username).Count(&count)
	if count > 0 {
		return errors.New("用户名已存在")
	}

	user := models.User{Username: username, Password: password}
	if err := database.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// 登录用户
func AuthenticateUser(username, password string) (*models.User, error) {
	var user models.User
	result := database.DB.Where("username = ? AND password = ?", username, password).First(&user)
	if result.Error != nil {
		return nil, errors.New("用户名或密码错误")
	}

	// 登录成功，存入 Redis
	_, err := SetUserSession(user.ID, user.Username)
	if err != nil {
		return nil, errors.New("Redis 会话存储失败")
	}

	return &user, nil
}

// 退出登录
func LogoutUser(userID uint) error {
	return DeleteUserSession(userID)
}
