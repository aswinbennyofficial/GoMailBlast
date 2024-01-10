package util

import(
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"net/smtp"
)

func SendBulkEmail(userList []User){

	// Loading environment variables
	err := godotenv.Load(".env")
	if err != nil {
        log.Fatalf("Error loading environment variables file")
    }

	// SMTP server Credentials from .env file
	SMTP_USERNAME := os.Getenv("SMTP_USERNAME")
	SMTP_PASSWORD := os.Getenv("SMTP_PASSWORD")
	SMTP_HOST :=os.Getenv("SMTP_HOST")
	FROM_EMAIL :=os.Getenv("FROM_EMAIL")
	SMTP_PORT :=os.Getenv("SMTP_PORT")
	REPLY_TO :=os.Getenv("REPLY_TO")
	
	log.Println("SMTP CREDS init ",SMTP_USERNAME, " ", SMTP_PASSWORD," ",SMTP_HOST )
	
	// Setup authentication variable
	auth:=smtp.PlainAuth("",SMTP_USERNAME,SMTP_PASSWORD,SMTP_HOST)

	

	if REPLY_TO==""{
		REPLY_TO=FROM_EMAIL
	}


	for _, user := range userList {
		// mail
		subject:="Test Golang Email Sender"
		body:="<html><body><h1>Hi "+user.Name+" ,</h1> <br>this is an HTML-rich email template!</body></html>"
		
		

		var msg []byte
		//For basic text
		// msg = []byte(
		// 	"Reply-To: "+reply_to+"\r\n"+
		// 	"Subject: "+subject+"\r\n" +
		// 	"\r\n" +
		// 	body+"\r\n")

		//For rich html support
		msg = []byte(
			"From: "+FROM_EMAIL+"\r\n"+
			"Reply-To: " + REPLY_TO + "\r\n" +
				"Subject: " + subject + "\r\n" +
				"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n" +
				"\r\n" +
				body + "\r\n")

		recieverEmail := []string{user.Email} 
		
		// send the mail
		err = smtp.SendMail(SMTP_HOST+":"+SMTP_PORT, auth, FROM_EMAIL, recieverEmail, msg)

		// handling the errors
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}

	
	fmt.Println("Successfully sent mail to all user in the list")

	

}