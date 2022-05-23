// Package defines usecases for project.
package usecases

import (
	"errors"
	"os"

	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/tlacuilose/transaction-reporter/domain/entities"
)

// SendSummaryByEmail sends a summary using email service.
func SendSummaryByEmail(summary entities.Summary, email string) error {
	from := mail.NewEmail("Jose Tlacuilo", "pammicc@tutanota.com")
	subject := "Your transaction summary is here"
	to := mail.NewEmail("Account", email)

	plainTextContent := summary.PrintHTML()
	htmlContent := summary.PrintHTML()

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	apiKey := os.Getenv("SENDGRID_API_KEY")
	if apiKey == "" {
		return errors.New("Unable to contact email service.")
	}

	client := sendgrid.NewSendClient(apiKey)
	_, err := client.Send(message)
	if err != nil {
		return err
	}

	return nil
}
