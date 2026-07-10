package services

import (
	"url-shortener-api/models"
	"url-shortener-api/repository"
	"url-shortener-api/utils"

	"errors"

	"gorm.io/gorm"
)

func RegisterUser(req models.SignupRequest) error {

	_, err := repository.GetUserByEmail(req.Email)

	if err == nil {
		return errors.New("email already exists")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return err
	}

	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}
	return repository.CreateUser(&user)
}
