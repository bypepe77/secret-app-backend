package database

import (
	"fmt"

	"github.com/bypepe77/secret-app-backend/internal/models"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DatabaseConnectionEnt() *gorm.DB {
	dsn := "root:rootpassword@tcp(127.0.0.1:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println("cause err", err)
	if err != nil {
		fmt.Println("err", err)
	}
	db.AutoMigrate(&models.User{}, &models.Confession{})

	return db
}
