package common

import (
	"github.com/goweb/config"
	"fmt"
	"net/smtp"
)

func SendMail(to string, title string, content string) error {


	host := config.MailConfig.Host
	//port := config.MailConfig.Port
	from := config.MailConfig.From
	pwd := config.MailConfig.Pwd
	mail := config.MailConfig.Email

	header := make(map[string]string)
	header["from"] = from
	header["to"] = mail
	header["subject"] = title
	header["Content-Type"] = "text/html;charset=utf-8"

	message := ""

	for k, v := range  header {
		message += fmt.Sprintf("%s :%s\r\n", k,v)
	}
	smtp.PlainAuth("",mail, pwd,host)





	return nil

}
