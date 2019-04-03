package config

import "log"

type serverConfig struct {
	Env string
	LogFile string
}


var ServerConfig serverConfig


func initServer(){

	log.Printf("init server now...")
	ServerConfig = serverConfig{
		Env:"dev",
		LogFile:"",
	}
}

func init() {
	initServer()
}