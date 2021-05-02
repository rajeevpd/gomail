package main

import (
	"log"
	"gopkg.in/gomail.v2"
	"fmt"
)

type SMTPClient struct {
	SMTPServer   string `json:"smtp_server"`
	SMTPUsername string `json:"smtp_username"`
	SMTPPassword string `json:"smtp_password"`
	SMTPPort     int    `json:"smtp_port"`
	SendTo       string `json:"send_to"`
}

func NewSMTPClient() *SMTPClient {
	return &SMTPClient{
		SMTPServer:   "smtp.gmail.com",
		SMTPUsername: "pleasechangethis@gmail.com",
		SMTPPassword: "passpass",
		SMTPPort:     587,
		SendTo:       "sendtoid@gmail.com",
	}
}

func (c *SMTPClient) Send() error {
	log.Print("Sending notification to: ", c.SendTo)
	m := gomail.NewMessage()

	m.SetHeader("From", c.SMTPUsername)
	m.SetHeader("To", c.SendTo)
	m.SetHeader("Subject", "go-mail")
	m.SetBody("text/html", "This is sample email body")

	d := gomail.NewDialer(
		c.SMTPServer,
		c.SMTPPort,
		c.SMTPUsername,
		c.SMTPPassword)

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func main() {
	mailClient := NewSMTPClient()
	err := mailClient.Send()

	if err != nil {
		fmt.Println("Error in sending email", err)
	}
}
