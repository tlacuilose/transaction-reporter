// Package defines an email service.
package email_service

import "testing"

// Test that sending an email is possible.
func TestSendEmail(t *testing.T) {
	if err := SendMail("Testing...", "test@email.com"); err != nil {
		t.Fatal(err)
	}
}
