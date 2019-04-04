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

type mailConfig struct {
	Host string
	Port int
	Email string
	From string
	ContentType string
	Pwd string
}



var ServerConfig serverConfig

var MysqlConfig dbConfig

var MailConfig mailConfig

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
		Dialect:"mysql",
	}
	}

func initMail()  {

}

func init() {
	initServer()
	initMysql()
	initMail()
}