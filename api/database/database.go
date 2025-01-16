package database

import (
	"crossx/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDB() {
	host := "localhost"
	user := "gorm"
	password := "gorm"
	dbname := "crossx"
	port := "5432"

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host,
		user,
		password,
		dbname,
		port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Post{})
	db.AutoMigrate(&models.Comment{})
}
