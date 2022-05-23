// Package defines the server that listens for requests.
package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/tlacuilose/transaction-reporter/domain/usecases"
)

// Start a server connected to a lambda proxy
func StartLambdaServer() {
	lambda.Start(handler)
}

// Handler function for the proxy request.
func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	account := request.QueryStringParameters["account"]
	email := request.QueryStringParameters["email"]
	if email != "" {
		sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))

		storeFile := fmt.Sprintf("%s.csv", account)
		transactions, err := usecases.ReadFromBucket("store-transactions", storeFile, sess)
		if err != nil {
			return events.APIGatewayProxyResponse{
				Body:       errors.New("Account not found").Error(),
				StatusCode: http.StatusBadRequest,
			}, nil
		}

		summary, err := usecases.CreateSummaryFromTransactions(transactions, account)
		if err != nil {
			return events.APIGatewayProxyResponse{
				Body:       errors.New("Failed to construct summary").Error(),
				StatusCode: http.StatusInternalServerError,
			}, nil
		}

		err = usecases.SendSummaryByEmail(summary, email)
		if err != nil {
			return events.APIGatewayProxyResponse{
				Body:       errors.New("Could not contact email service").Error(),
				StatusCode: http.StatusInternalServerError,
			}, nil

		}

		return events.APIGatewayProxyResponse{
			Body:       "Email sent! to: " + email,
			StatusCode: http.StatusOK,
		}, nil

	}

	return events.APIGatewayProxyResponse{
		Body:       "Bad POST request",
		StatusCode: http.StatusBadRequest,
	}, nil
}
