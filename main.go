package main

import (
	"flag"
	"time"

	"github.com/JisungPark0319/go-gin-boilerplate/auth"
	"github.com/JisungPark0319/go-gin-boilerplate/config"
	"github.com/JisungPark0319/go-gin-boilerplate/database"
	"github.com/JisungPark0319/go-gin-boilerplate/models"
	"github.com/JisungPark0319/go-gin-boilerplate/router"
	"github.com/gin-gonic/gin"
)

var (
	configPath string
)

func main() {
	var err error

	flag.StringVar(&configPath, "config", "./config.yaml", "Specify config file path")
	flag.Parse()

	cfg, err := config.LoadConfig(configPath, "GIN")
	if err != nil {
		panic(err)
	}
	engine := gin.Default()

	database.Init(cfg.DatabaseConfig)
	err = models.AutoMigrate()
	if err != nil {
		panic(err)
	}
	auth.New(cfg.AuthConfig)
	auth.Get().SetExpire(time.Minute*10, time.Hour*1)

	// engine.Use(gin.Logger())

	router.Set(engine)
	engine.Run()
}
