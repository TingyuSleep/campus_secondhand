package services

import (
	"campus_secondhand/models"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func Register(username, password string) (*models.User, error) {
	existingUser, err := models.GetUserByUsername(username)
	if err == nil && existingUser != nil {
		return nil, errors.New("用户名已存在")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: username,
		Password: string(hashedPassword),
	}

	err = models.AddUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func Login(username, password string) (*models.User, error) {
	user, err := models.GetUserByUsername(username)
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	return user, nil
}
