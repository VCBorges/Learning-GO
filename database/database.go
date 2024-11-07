package database

import (
	"project_name/users"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("dev.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(
		users.User{},
	)
	return db, nil
}
