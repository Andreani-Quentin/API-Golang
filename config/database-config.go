package config

import (
	"festApp/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

//SetupDatabaseConnection is creating a new connection to our database
func SetupDatabaseConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("festApp.db"), &gorm.Config{})
	if err != nil {
		panic("Erreur de connection à la base de donnée !")
	}

	err = db.AutoMigrate(&entity.Event{}, &entity.Artist{}, &entity.User{}, &entity.Event{})
	if err != nil {
		panic("Erreur de migration à la base de donnée !")
	}

	return db
}

//CloseDatabaseConnection method is closing a connection between your app and your db
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}

