package config

import (
	"github.com/marcospsw/first-project-go/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger()
	// Create DB and Connect
	db, err := gorm.Open(sqlite.Open("./db/main.db"), &gorm.Config{})
	if err != nil {
		logger.Errf("Error on connection of database: %v", err)
		return nil, err
	}
	err = db.AutoMigrate(&schemas.Opening{})
	return db, nil
}
