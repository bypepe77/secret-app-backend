package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"not null;unique"`
	Email    string
	Password string
	IsBanned bool `gorm:"default:false"`
	IsActive bool `gorm:"default:true"`
}
