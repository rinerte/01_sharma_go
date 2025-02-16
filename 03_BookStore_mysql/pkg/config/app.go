package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "host=localhost user=postgres password=postgres dbname=go_bookstore port=5432 sslmode=disable TimeZone=UTC"
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
