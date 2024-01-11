package util

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
	"github.com/skip2/go-qrcode"
	//"strings"
)

func SendBulkEmail(userList []User) {
	// Loading environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading environment variables file")
	}

	// SMTP server Credentials from .env file
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
		qrCode, err := generateQRCode(user.Name)
		if err != nil {
			log.Println("Error generating QR code:", err)
			continue
		}

		// Convert QR code to PNG
		pngData, err := convertQRCodeToPNG(qrCode)
		if err != nil {
			log.Println("Error converting QR code to PNG:", err)
			continue
		}

		// Create email body
		subject := "Test Golang Email Sender"
		body := fmt.Sprintf("<html><body><h1>Hi %s,</h1> <br>this is an HTML-rich email template!<br><img src=\"cid:qrcode\"></body></html>", user.Name)

		// Create MIME email with embedded image
		msg := []byte(
			"From: " + FROM_EMAIL + "\r\n" +
				"Reply-To: " + REPLY_TO + "\r\n" +
				"Subject: " + subject + "\r\n" +
				"MIME-version: 1.0;\nContent-Type: multipart/related; boundary=\"related_boundary\";\r\n" +
				"\r\n" +
				"--related_boundary\r\n" +
				"Content-Type: text/html; charset=\"UTF-8\";\r\n" +
				"\r\n" +
				body + "\r\n" +
				"--related_boundary\r\n" +
				"Content-Type: image/png; name=\"qrcode.png\"\r\n" +
				"Content-Disposition: inline; filename=\"qrcode.png\"\r\n" +
				"Content-ID: <qrcode>\r\n" +
				"Content-Transfer-Encoding: base64\r\n" +
				"\r\n" +
				pngData + "\r\n" +
				"--related_boundary--")

		// Recipient email
		recieverEmail := []string{user.Email}

		// Send the mail
		err = smtp.SendMail(SMTP_HOST+":"+SMTP_PORT, auth, FROM_EMAIL, recieverEmail, msg)

		// Handle errors
		if err != nil {
			log.Println(err)
			continue
		}
	}

	fmt.Println("Successfully sent mail to all users in the list")
}

func generateQRCode(data string) (*qrcode.QRCode, error) {
	qrCode, err := qrcode.New(data, qrcode.Medium)
	if err != nil {
		return nil, err
	}
	return qrCode, nil
}

func convertQRCodeToPNG(qrCode *qrcode.QRCode) (string, error) {
	// Create a new PNG image from the QR code
	pngData, err := qrCode.PNG(256)
	if err != nil {
		return "", err
	}

	// Encode the PNG image to base64
	pngBase64 := base64.StdEncoding.EncodeToString(pngData)

	return pngBase64, nil
}
