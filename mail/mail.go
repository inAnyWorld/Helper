package mail

import (
	"gopkg.in/gomail.v2"
	"log"
	"strings"
)

// Send 发邮件
func Send(to string, subject string, body string) error {
	mClient := gomail.NewMessage()
	mClient.SetHeader("From", "username")
	mailTo := strings.Split(to, ",")
	mClient.SetHeader("To", mailTo ...)
	mClient.SetHeader("Subject", subject)
	mClient.SetBody("text/html", body)
	newDialer := gomail.NewDialer("host", 587, "username", "password")
	err := newDialer.DialAndSend(mClient)
	if err != nil {
		log.Println(err)
	}
	return err
}
