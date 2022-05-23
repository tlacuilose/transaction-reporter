// Package defines an email service.
package email_service

import "fmt"

// DummySendEmail has not email sender but prints the
func DummySendMail(body string, email string) error {
	fmt.Printf("Mail sent! to: %s\n", email)
	return nil
}
