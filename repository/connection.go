package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(DSN string) (*gorm.DB, error) {
	// connect ke database, DSN ambil dari .env
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
