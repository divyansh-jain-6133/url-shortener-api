package models

import "gorm.io/gorm"

type URL struct {
	gorm.Model

	LongURL  string `gorm:"not null"`
	ShortURL string `gorm:"unique;not null"`

	Clicks uint `gorm:"default:0"`

	UserID uint
}
