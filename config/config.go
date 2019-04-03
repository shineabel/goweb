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

var MongoConfig dbConfig


func initServer(){

	log.Printf("init server now...")
	ServerConfig = serverConfig{
		Env:"dev",
		LogFile:"",
	}
}

func initMongo()  {
	log.Println("init mongo...")

	MongoConfig = dbConfig{
		URL:"",
		UserName:"",
		Password:"",
		Host:"",
		Port:"",
		Dialect:"",
		Database:"",
	}

}

func init() {
	initServer()
	initMongo()
}