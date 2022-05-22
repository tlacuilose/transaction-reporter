// Package defines usecases for project.
package usecases

import (
	"github.com/tlacuilose/transaction-reporter/data/services/email_service"
	"github.com/tlacuilose/transaction-reporter/domain/entities"
)

// SendSummaryByEmail sends a summary using email service.
func SendSummaryByEmail(summary entities.Summary, email string) error {
	return email_service.SendMail(summary.PrintHTML(), email)
}
