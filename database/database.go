package database

import (
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	database, err := getDatabaseFactory("postgres")
	if err != nil {
		panic(err)
	}

	err = database.connect("localhost", "35432", "postgres", "postgres", "postgres")
	if err != nil {
		panic(err)
	}

	db = database.getDB()
}

func GetDB() *gorm.DB {
	return db
}

func AutoMigrate(models []interface{}) error {
	err := db.AutoMigrate(models...)
	return err
}
