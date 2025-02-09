package models

import (
	"campus_secondhand/database"
)

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"type:varchar(255);uniqueIndex;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
}

func Migrate() {
	err := database.DB.AutoMigrate(&User{})
	if err != nil {
		panic("数据库迁移失败")
	}
}

func AddUser(user *User) error {
	result := database.DB.Create(user)
	return result.Error
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	result := database.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
