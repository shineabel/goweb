package db

import (
	"github.com/goweb/config"
	"fmt"
	"github.com/jinzhu/gorm"
)
var DB *gorm.DB

func initDB() {

	db , err := gorm.Open(config.MongoConfig.Dialect, config.MongoConfig.URL)
	if err != nil {
		fmt.Errorf("init db error:", err)
		return
	}

	DB = db

}

