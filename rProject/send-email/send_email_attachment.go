package sendemail

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"strings"
)

type Mail struct {
	Sender      string
	To          []string
	Subject     string
	Body        string
	Attachments string
}

func SendEmailWithAttachment(email string, message string) {

	namaFile := Mail{}
	namaFile.Attachments = "send_email.go"

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
	filePath := "./send-email/" + fileName
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Panic(err)
	}

	return data
}
