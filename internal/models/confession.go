package models

import (
	"gorm.io/gorm"
)

type Confession struct {
	gorm.Model
	Content     string
	IsDestroyed bool `gorm:"default:true"`
	User        *User
	UserID      int `gorm:"foreignkey:UserID"`
}
