package models

import "github.com/JisungPark0319/go-gin-boilerplate/database"

func AutoMigrate() error {
	var models = []interface{}{
		&User{},
	}

	err := database.AutoMigrate(models)
	if err != nil {
		return err
	}
	return nil
}
