package main

import (
	"fmt"
	"time"

	"github.com/JisungPark0319/go-gin-boilerplate/auth"
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
	auth.New("accessSecret", "refreshSecret")
	auth.Get().SetExpire(time.Minute*10, time.Hour*1)

	router.Run()
}
