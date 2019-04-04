package common

import (
	"github.com/goweb/config"
	"fmt"
	"net/smtp"
	"crypto/tls"
	"net"
)

func SendMail(to string, title string, content string) error {


	host := config.MailConfig.Host
	port := config.MailConfig.Port
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
	auth := smtp.PlainAuth("",mail, pwd,host)


	err := sendEmailUseTLS(fmt.Sprintf("%s:%d", host, port), auth,mail,[]string{to},content)
	return err

}

func sendEmailUseTLS(addr string, auth smtp.Auth, from string, to []string, message string) error  {

	client, err := createSMTPClient(addr)
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	defer client.Close()

	if auth != nil {
		if ok , _ := client.Extension("AUTH"); ok {
			err := client.Auth(auth)
			if err != nil {
				fmt.Printf(err.Error())
				return err
			}
		}
	}

	if err := client.Mail(from); err != nil {
		return err
	}



	 for _, addr := range  to {
	 	if err := client.Rcpt(addr); err != nil {
	 		return err
		}
	 }


	wc, err := client.Data()
	if err != nil {
		return err
	}

	_ , err = wc.Write([]byte(message))
	if err != nil {
		return err
	}
	err = wc.Close()
	if err != nil {
		return err
	}
	return client.Quit()
}

func createSMTPClient(addr string)(c *smtp.Client,  err error)  {

	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		fmt.Printf(err.Error())
		return nil, err
	}
	host, _, _ := net.SplitHostPort(addr)

	return smtp.NewClient(conn, host)
}