package db

import (
	"github.com/goweb/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
var DB *gorm.DB

func init() {
	fmt.Println("init db, url:", config.MysqlConfig.URL)

	db , err := gorm.Open(config.MysqlConfig.Dialect,config.MysqlConfig.URL)
	if err != nil {
		fmt.Println("init db error:", err)
		return
	}

	fmt.Println("init mysql OK")
	DB = db

}

