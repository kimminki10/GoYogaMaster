package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Migrate automigrates models using ORM
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{}, &Pose{}, &Content{})

	db.Model(&Pose{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	db.Model(&Content{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	fmt.Println("Auto Migration has been processed")
}
