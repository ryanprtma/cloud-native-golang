package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"io/ioutil"
	"time"

	"fmt"
	"log"
	"net/http"

	"rProject/config"
	"rProject/utils"

	"encoding/json"
	"net/smtp"
	"rProject/models"
	"rProject/req"
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

	fmt.Println("Server Run on :7000")

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
	// return
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
	// return
}

func sendBulkEmails(bulkEmails []string, message string) {
	for i, email := range bulkEmails {
		SendEmailWithAttachment(email, message)
		if i == 3 {
			time.Sleep(60 * time.Second)
			continue
		}
		// sendemail.SendSingleEmail(email, message)
	}
}

func SendEmailWithAttachment(email string, message string) {

	namaFile := Mail{}
	namaFile.Attachments = "main.go"

	sender := "test@example.com"

	to := []string{
		email,
	}

	user := "4ddb3336739147"
	password := "2c24a017a3e258"

	subject := "testing mail with attachment"
	body := message

	request := Mail{
		Sender:  sender,
		To:      to,
		Subject: subject,
		Body:    body,
	}

	addr := "smtp.mailtrap.io:2525"
	host := "smtp.mailtrap.io"

	data := BuildMail(request, namaFile.Attachments)
	auth := smtp.PlainAuth("", user, password, host)
	err := smtp.SendMail(addr, auth, sender, to, data)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Email sent successfully")

}

func BuildMail(mail Mail, namafile string) []byte {

	var buf bytes.Buffer

	buf.WriteString(fmt.Sprintf("From: %s\r\n", mail.Sender))
	buf.WriteString(fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";")))
	buf.WriteString(fmt.Sprintf("Subject: %s\r\n", mail.Subject))

	boundary := "my-boundary-779"
	buf.WriteString("MIME-Version: 1.0\r\n")
	buf.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\n",
		boundary))

	buf.WriteString(fmt.Sprintf("\r\n--%s\r\n", boundary))
	buf.WriteString("Content-Type: text/plain; charset=\"utf-8\"\r\n")
	buf.WriteString(fmt.Sprintf("\r\n%s", mail.Body))
	data := readFile(namafile)

	b := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	buf.WriteString(fmt.Sprintf("\r\n--%s\r\n", boundary))
	buf.WriteString(fmt.Sprintf("Content-Type: %s\n", http.DetectContentType(b)))
	buf.WriteString("Content-Transfer-Encoding: base64\r\n")
	buf.WriteString("Content-Disposition: attachment; filename=" + namafile + "\r\n")
	buf.WriteString("Content-ID: <" + namafile + ">\r\n\r\n")

	base64.StdEncoding.Encode(b, data)
	buf.Write(b)
	buf.WriteString(fmt.Sprintf("\r\n--%s", boundary))

	buf.WriteString("--")

	return buf.Bytes()
}

func readFile(fileName string) []byte {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
