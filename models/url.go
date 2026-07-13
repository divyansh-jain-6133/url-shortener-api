package models

import "gorm.io/gorm"

type URL struct {
	gorm.Model

	LongURL  string `gorm:"column:long_url;not null"`
	ShortURL string `gorm:"column:short_url;unique;not null"`
	Clicks   uint   `gorm:"default:0"`

	UserID uint `gorm:"column:user_id"`
}
