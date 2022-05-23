// Package defines an email service.
package email_service

import "testing"

// Test that sending an email is possible.
func TestDummySendEmail(t *testing.T) {
	if err := DummySendMail("Testing...", "test@email.com"); err != nil {
		t.Fatal(err)
	}
}
