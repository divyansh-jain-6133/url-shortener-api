package services

import (
	"url-shortener-api/models"
	"url-shortener-api/repository"
	"url-shortener-api/utils"
)

func CreateShortURL(req models.ShortenRequest, userID uint) (*models.URL, error) {

	shortCode := utils.GenerateShortCode(6)

	url := models.URL{
		LongURL:  req.URL,
		ShortURL:   shortCode,
		UserID:      userID,
	}

	err := repository.CreateURL(&url)

	if err != nil {
		return nil, err
	}

	return &url, nil
}

func GetURLByCode(code string) (*models.URL, error) {
	return repository.GetURLByCode(code)
}
