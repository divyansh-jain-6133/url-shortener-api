package repository

import (
	"url-shortener-api/config"
	"url-shortener-api/models"
)

func CreateURL(url *models.URL) error {
	return config.DB.Create(url).Error
}

func GetURLByCode(code string) (*models.URL, error) {
	var url models.URL

	err := config.DB.Where("short_url = ?", code).First(&url).Error

	return &url, err
}
