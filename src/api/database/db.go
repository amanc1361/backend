package database

import (
	"back-account/src/api/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {

	db, error := gorm.Open(mysql.Open(config.DSN), &gorm.Config{})

	if error != nil {
		return nil, error
	}

	return db, error
}
