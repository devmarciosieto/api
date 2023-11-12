package email

import (
	"fmt"
	"github.com/devmarciosieto/api/internal/domain/campaign"
	"gopkg.in/gomail.v2"
	"os"
)

func SendEmail(campaign *campaign.Campaign) error {
	fmt.Println("Sending email...")

	var emails []string

	for _, contact := range campaign.Contacts {
		emails = append(emails, contact.Email)
	}

	d := gomail.NewDialer(os.Getenv("EMAIL_SMTP"), 587, os.Getenv("EMAIL_USER"), os.Getenv("EMAIL_PASSWORD"))

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_FROM"))
	m.SetHeader("To", emails...)
	m.SetHeader("Subject", campaign.Name)
	m.SetBody("text/html", campaign.Content)

	return d.DialAndSend(m)
}
