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
	LikesCount  int32       `gorm:"default:0"`
	LikesList   []*Like     `gorm:"many2many:confession_likes;"`
}

type Category struct {
	gorm.Model
	Name string
}

type Like struct {
	gorm.Model
	User         *User
	UserID       int `gorm:"foreignkey:UserID"`
	Confession   *Confession
	ConfessionID int `gorm:"foreignkey:ConfessionID"`
}
