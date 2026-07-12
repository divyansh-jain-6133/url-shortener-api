package models

import "gorm.io/gorm"

type URL struct {
	gorm.Model

	LongURL  string `gorm:"not null"`
	ShortURL string `gorm:"unique;not null"`

	UserID uint
}
