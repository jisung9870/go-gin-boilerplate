package main

import (
	"fmt"

	"github.com/JisungPark0319/go-gin-boilerplate/database"
	"github.com/JisungPark0319/go-gin-boilerplate/models"
	"github.com/JisungPark0319/go-gin-boilerplate/router"
)

func main() {
	database.Init()
	err := models.AutoMigrate()
	if err != nil {
		fmt.Println(err)
		return
	}

	router.Run()
}
