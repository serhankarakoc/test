package notifiers

import (
	"fmt"
	"net/smtp"
	"os"
)

type EmailNotifier struct{}

func (n EmailNotifier) Send(content map[string]string) {
	from := os.Getenv("MAIL_FROM_ADDRESS")
	pass := os.Getenv("MAIL_PASSWORD")
	to := content["email"]
	subject := content["title"]
	body := content["content"]

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	err := smtp.SendMail(os.Getenv("MAIL_HOST")+":"+os.Getenv("MAIL_PORT"),
		smtp.PlainAuth("", from, pass, os.Getenv("MAIL_HOST")),
		from, []string{to}, []byte(msg))

	if err != nil {
		fmt.Printf("Error while sending email: %v", err)
		return
	}

	fmt.Printf("Email sent successfully to %s\n", to)
}
