package config

import (
	"log"
)

type serverConfig struct {
	Env string
	LogFile string
}

type dbConfig struct {
	URL string
	UserName string
	Password string
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
	}

}

func init() {
	initServer()
	initMongo()
}