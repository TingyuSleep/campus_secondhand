package services

import (
	"campus_secondhand/database"
	"fmt"
	"time"
)

// 设置用户 Session，登录成功后存入 Redis
func SetUserSession(userID uint, username string) (string, error) {
	sessionKey := fmt.Sprintf("session:user:%d", userID)
	err := database.Rdb.Set(database.Ctx, sessionKey, username, time.Hour*24).Err()
	if err != nil {
		return "", err
	}
	return sessionKey, nil
}

// 获取用户 Session
func GetUserSession(userID uint) (string, error) {
	sessionKey := fmt.Sprintf("session:user:%d", userID)
	username, err := database.Rdb.Get(database.Ctx, sessionKey).Result()
	if err != nil {
		return "", err
	}
	return username, nil
}

// 删除用户 Session (退出登录)
func DeleteUserSession(userID uint) error {
	sessionKey := fmt.Sprintf("session:user:%d", userID)
	err := database.Rdb.Del(database.Ctx, sessionKey).Err()
	return err
}
