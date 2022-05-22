// Package defines an email service.
package email_service

import "fmt"

// SendMail sends an email.
func SendMail(body string, email string) error {
	fmt.Printf("Mail sent! to: %s\n", email)
	return nil
}
