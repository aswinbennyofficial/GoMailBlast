package util

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
	qrcode "github.com/skip2/go-qrcode"
)

func SendBulkEmail(userList []User) {
	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading environment variables file")
	}

	// SMTP server credentials from .env file
	SMTP_USERNAME := os.Getenv("SMTP_USERNAME")
	SMTP_PASSWORD := os.Getenv("SMTP_PASSWORD")
	SMTP_HOST := os.Getenv("SMTP_HOST")
	FROM_EMAIL := os.Getenv("FROM_EMAIL")
	SMTP_PORT := os.Getenv("SMTP_PORT")
	REPLY_TO := os.Getenv("REPLY_TO")

	log.Println("SMTP CREDS init ", SMTP_USERNAME, " ", SMTP_PASSWORD, " ", SMTP_HOST)

	// Setup authentication variable
	auth := smtp.PlainAuth("", SMTP_USERNAME, SMTP_PASSWORD, SMTP_HOST)

	if REPLY_TO == "" {
		REPLY_TO = FROM_EMAIL
	}

	for _, user := range userList {
		// Generate QR code
		qrCode, err := generateQRCode("iudshgiugigfisigfisiugiug")
		if err != nil {
			log.Println("Error generating QR code:", err)
			continue
		}

		// Mail
		subject := "Test Golang Email Sender 2"
		body := "<html><body><h1>Hi " + user.Name + "</h1><br>This is an HTML-rich email template!<br>Join fast</body></html>"

		var msg []byte
		msg = []byte(
			"From: " + FROM_EMAIL + "\r\n" +
				"Reply-To: " + REPLY_TO + "\r\n" +
				"Subject: " + subject + "\r\n" +
				"MIME-version: 1.0;\nContent-Type: multipart/mixed; boundary=foo;\r\n\r\n" +
				"--foo\r\n" +
				"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
				"\r\n" +
				body + "\r\n" +
				"--foo\r\n" +
				"Content-Type: image/png\r\n" +
				"Content-Transfer-Encoding: base64\r\n" +
				"Content-Disposition: attachment; filename=\"qr-code.png\"\r\n" +
				"\r\n" +
				qrCode + "\r\n" +
				"--foo--\r\n")

		recieverEmail := []string{user.Email}

		// Send the mail
		err = smtp.SendMail(SMTP_HOST+":"+SMTP_PORT, auth, FROM_EMAIL, recieverEmail, msg)

		// Handling the errors
		if err != nil {
			log.Println(err)
			continue
		}
	}

	fmt.Println("Successfully sent mail to all users in the list")
}

func generateQRCode(data string) (string, error) {
	// Generate QR code and return it as a base64-encoded string
	qrCode, err := qrcode.Encode(data, qrcode.Medium, 256)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(qrCode), nil
}
