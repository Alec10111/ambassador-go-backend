package database

import (
	"ambassador/src/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	dsn := "host=amb-db user=amb password=amb dbname=amb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database.")
	}
}

func AutoMigrate() {
	DB.AutoMigrate((models.User{}))
}
