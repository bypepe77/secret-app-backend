package models

import (
	"gorm.io/gorm"
)

type Confession struct {
	gorm.Model
	Content     string
	IsDestroyed bool `gorm:"default:true"`
	User        *User
	UserID      int         `gorm:"foreignkey:UserID"`
	Categories  []*Category `gorm:"many2many:confession_categories;"`
}

type Category struct {
	gorm.Model
	Name string
}
