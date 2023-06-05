package models

import (
	"gorm.io/gorm"
)

type SharedProduct struct {
	
	gorm.Model

	UserID     int    `json:"userId"`
	ProductID  int    `json:"productId"`
	SocialMedia string `json:"socialMedia"`
	Link     string `gorm:"not null"`
}




