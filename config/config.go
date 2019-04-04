package config

import (
	"log"
)

type serverConfig struct {
	Env string
	LogFile string
}

type dbConfig struct {
	Host string
	Port string
	URL string
	UserName string
	Password string
	Dialect string
	Database string
}



var ServerConfig serverConfig

var MysqlConfig dbConfig


func initServer(){

	log.Println("init server now...")
	ServerConfig = serverConfig{
		Env:"dev",
		LogFile:"./goweb.log",
	}
}

func initMysql()  {
	log.Println("init mysql...")

	MysqlConfig = dbConfig{
		URL:"root:shine@tcp(localhost:3306)/test?charset=utf8",
		UserName:"root",
		Password:"shine",
		Host:"localhost",
		Port:"3306",
		Dialect:"mysql",
		Database:"test",
	}

}

func init() {
	initServer()
	initMysql()
}