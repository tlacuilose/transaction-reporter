// Package defines usecases for project.
package usecases

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"

	"github.com/tlacuilose/transaction-reporter/domain/entities"
)

const (
	Sender = "pammicc@tutanota.com"

	Subject = "Your transaction summary from Test"

	CharSet = "UTF-8"
)

// SendSummaryByEmail sends a summary using email service.
func SendSummaryWithSES(summary entities.Summary, email string, sess *session.Session) error {
	svc := ses.New(sess)

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			CcAddresses: []*string{},
			ToAddresses: []*string{
				aws.String(email),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(summary.PrintHTML()),
				},
				Text: &ses.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(summary.PrintHTML()),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(Subject),
			},
		},
		Source: aws.String(Sender),
	}

	_, err := svc.SendEmail(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			return aerr
		} else {
			return err
		}
	}
	return nil
}
