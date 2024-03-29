package database

import (
	"fmt"
	"os"

	"mingi/goyoma/database/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // configures mysql driver
)

// Initialize initializes the database
func Initialize() (*gorm.DB, error) {
	dbConfig := os.Getenv("DB_CONFIG")
	db, err := gorm.Open("mysql", dbConfig)
	db.LogMode(true)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")
	models.Migrate(db)
	return db, err
}
