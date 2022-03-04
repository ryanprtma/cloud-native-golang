package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"rProject/config"
	"rProject/utils"

	"encoding/json"
	"rProject/models"
	"rProject/req"
	"strings"
)

const CONFIG_SMTP_HOST = "smtp.mailtrap.io"
const CONFIG_SMTP_PORT = 587
const CONFIG_SENDER_NAME = "test <testSend@gmail.com>"
const CONFIG_AUTH_EMAIL = "4ddb3336739147"
const CONFIG_AUTH_PASSWORD = "2c24a017a3e258"

func main() {

	db, e := config.MySQL()

	if e != nil {
		log.Fatal(e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}

	fmt.Println("Success Connect to database")

	http.HandleFunc("/reqemails", GetReqEmails)
	http.HandleFunc("/reqemails/create", PostReqEmails)

	err := http.ListenAndServe(":7000", nil)

	if err != nil {
		log.Fatal(err)
	}
}

// GetReq
func GetReqEmails(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		ctx, cancel := context.WithCancel(context.Background())

		defer cancel()

		reqs, err := req.GetAll(ctx)

		if err != nil {
			fmt.Println(err)
		}

		utils.ResponseJSON(w, reqs, http.StatusOK)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusNotFound)
	return
}

func PostReqEmails(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Content type must JSON Application / json", http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		var post models.Req

		if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
			utils.ResponseJSON(w, err, http.StatusBadRequest)
			return
		}

		if err := req.Insert(ctx, post); err != nil {
			utils.ResponseJSON(w, err, http.StatusInternalServerError)
			return
		}

		// emails := map[string]string{
		// 	"email1": post.Email,
		// 	"emailx": "testx@gmail.com",
		// 	"emaily": "testy@gmail.com",
		// }

		// emailEntries(post.Email)

		var emailsSlice []string
		entries := strings.Split(post.Email, ",")
		for _, email := range entries {
			emailsSlice = append(emailsSlice, email)
		}

		res := map[string]string{
			"status": "Email Sent Succesfully!",
			"email":  post.Email,
			"text":   post.Text,
		}

		// go sendExecute(post.Email, post.Text)

		go sendBulkEmails(emailsSlice, post.Text)

		utils.ResponseJSON(w, res, http.StatusCreated)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}

func sendBulkEmails(bulkEmails []string, message string) {
	for _, email := range bulkEmails {
		sendSingleEmail(email, message)
	}
}

func sendSingleEmail(from string, pesan string) {
	to := []string{from}
	cc := []string{"testing@gmail.com"}
	subject := "Test mail"
	message := pesan
	sendMail(to, cc, subject, message)

}

func sendMail(to []string, cc []string, subject, message string) {
	body := "From: " + CONFIG_SENDER_NAME + "\n" +
		"To: " + strings.Join(to, ", ") + "\n" +
		"Cc: " + strings.Join(cc, ", ") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD, CONFIG_SMTP_HOST)
	smtpAddr := fmt.Sprintf("%s:%d", CONFIG_SMTP_HOST, CONFIG_SMTP_PORT)

	smtp.SendMail(smtpAddr, auth, CONFIG_AUTH_EMAIL, append(to, cc...), []byte(body))
}
