package main

import (
	"context"
	"time"

	"fmt"
	"log"
	"net/http"

	"rProject/config"
	"rProject/utils"

	"encoding/json"
	"rProject/models"
	"rProject/req"
	sendemail "rProject/send-email"
	"strings"
)

type Mail struct {
	Sender      string
	To          []string
	Subject     string
	Body        string
	Attachments string
}

func main() {

	db, e := config.MySQL()

	if e != nil {
		log.Fatal(e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}

	fmt.Println("Server Run on :9000")

	http.HandleFunc("/reqemails", GetReqEmails)
	http.HandleFunc("/reqemails/create", PostReqEmails)

	err := http.ListenAndServe(":9000", nil)

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

		go sendBulkEmails(emailsSlice, post.Text)

		utils.ResponseJSON(w, res, http.StatusCreated)
		return
	}

	http.Error(w, "Tidak di ijinkan", http.StatusMethodNotAllowed)
	return
}

func sendBulkEmails(bulkEmails []string, message string) {
	for i, email := range bulkEmails {
		sendemail.SendEmailWithAttachment(email, message)
		if i == 3 {
			time.Sleep(60 * time.Second)
			continue
		}
		// sendemail.SendSingleEmail(email, message)
	}
}
