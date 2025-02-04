package database

import (
	"crossx/models"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Test de SetupDB (simulation avec SQLite)
func TestSetupDB(t *testing.T) {
	// On remplace la base PostgreSQL par SQLite en mémoire pour les tests
	var err error
	db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	// Vérifier que la connexion s'est bien établie
	assert.NoError(t, err, "La connexion à la base de données ne doit pas renvoyer d'erreur")
	assert.NotNil(t, db, "L'objet DB ne doit pas être nil")

	// Migration des modèles
	err = db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	assert.NoError(t, err, "L'auto-migration des modèles ne doit pas échouer")
}

// Test de GetDB
func TestGetDB(t *testing.T) {
	// Vérifier que GetDB retourne bien une instance valide
	dbInstance := GetDB()
	assert.NotNil(t, dbInstance, "GetDB ne doit pas retourner nil")
}
