package main

import (
	"fmt"
	"log"
	"net/smtp"

	// "rProject/models"
	"strings"
)

const CONFIG_SMTP_HOST = "smtp.mailtrap.io"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "test <testSend@gmail.com>"
const CONFIG_AUTH_EMAIL = "4ddb3336739147"
const CONFIG_AUTH_PASSWORD = "2c24a017a3e258"

func main() {
	// var data models.Req

	to := []string{"test@gmail.com", "another.email@gmail.com"}
	cc := []string{"testing@gmail.com"}
	subject := "Test mail"
	message := "test"

	err := sendMail(to, cc, subject, message)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")
}

func sendMail(to []string, cc []string, subject, message string) error {
	body := "From: " + CONFIG_SENDER_NAME + "\n" +
		"To: " + strings.Join(to, " ,") + "\n" +
		"Cc: " + strings.Join(cc, " ,") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD, CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

	err := smtp.SendMail(smtpAddr, auth, CONFIG_AUTH_EMAIL, append(to, cc...), []byte(body))
	if err != nil {
		return err
	}

	return nil
}
