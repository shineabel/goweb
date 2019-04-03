package main

import (
	_ "github.com/goweb/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goweb/config"
	"os"
	"log"
	"io"
	"github.com/goweb/router"
)
func main() {

	fmt.Printf("gin version %s:", gin.Version)

	if config.ServerConfig.Env != "dev" {
		gin.SetMode(gin.ReleaseMode)
		gin.DisableConsoleColor()
		logFile, err := os.OpenFile(config.ServerConfig.LogFile, os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0666)
		if err != nil {
			log.Fatal("open log file failed")
			os.Exit(-1)
		}
		gin.DefaultWriter = io.MultiWriter(logFile)

	}

	app := gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	router.Route(app)

	app.Run(":8081")

}
