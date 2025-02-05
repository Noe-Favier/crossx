package database

import (
	"crossx/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB // Variable globale pour stocker la connexion

func SetupDB() {
	host := "localhost"
	user := "gorm"
	password := "gorm"
	dbname := "crossx"
	port := "5432"

	var err error

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host,
		user,
		password,
		dbname,
		port,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Post{})
	db.AutoMigrate(&models.Comment{})
}

// GetDB retourne la connexion à la base de données
func GetDB() *gorm.DB {
	return db
}

// SetTestDB return une db pour les test
func SetTestDB(testDB *gorm.DB) {
	db = testDB
}
