package models

import "gorm.io/gorm"


type UserMediaToken struct {

	gorm.Model
  
  //  ID           uint      `gorm:"primary_key"`
    Username     string    `gorm:"not null"`
    AccessToken  string    `gorm:"not null"`
    TokenType    string    `gorm:"not null"`
    ExpiresIn    int       `gorm:"not null"`
    RefreshToken string    `gorm:"not null"`
    // CreatedAt    time.Time `gorm:"not null"`
    // UpdatedAt    time.Time `gorm:"not null"`
}

