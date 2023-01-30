package database

import (
	"github.com/JisungPark0319/go-gin-boilerplate/config"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(cfg config.DatabaseConfig) {
	database, err := getDatabaseFactory("postgres")
	if err != nil {
		panic(err)
	}

	err = database.connect(cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DatabaseName)
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
