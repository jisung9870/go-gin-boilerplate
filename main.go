package main

import (
	"flag"
	"os"
	"time"

	"github.com/JisungPark0319/go-gin-boilerplate/auth"
	"github.com/JisungPark0319/go-gin-boilerplate/config"
	"github.com/JisungPark0319/go-gin-boilerplate/database"
	"github.com/JisungPark0319/go-gin-boilerplate/models"
	"github.com/JisungPark0319/go-gin-boilerplate/router"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	configPath string
)

func main() {
	var err error

	flag.StringVar(&configPath, "config", "./config.yaml", "Specify config file path")
	flag.Parse()

	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	cfg, err := config.LoadConfig(configPath, "GIN")
	if err != nil {
		panic(err)
	}

	engine := gin.Default()

	engine.Use(gin.Recovery())

	database.Init(cfg.DatabaseConfig)
	err = models.AutoMigrate()
	if err != nil {
		panic(err)
	}
	auth.New(cfg.AuthConfig)
	auth.Get().SetExpire(time.Minute*10, time.Hour*1)

	router.Set(engine)
	engine.Run()
}
