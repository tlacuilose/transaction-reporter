// Package defines the server that listens for requests.
package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tlacuilose/transaction-reporter/domain/usecases"
)

// EchoServer uses echo to serve REST Api.
type EchoServer struct {
	e    *echo.Echo
	port uint
}

func New(port uint) *EchoServer {
	e := echo.New()
	return &EchoServer{e, port}
}

func (s *EchoServer) Start() {
	e := s.e
	e.POST("email/:account", emailSummary)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", s.port)))
}

func emailSummary(c echo.Context) error {
	account := c.Param("account")
	email := c.QueryParam("email")

	storeFile := fmt.Sprintf("./store/%s.csv", account)

	transactions, err := usecases.ReadFromLocalFile(storeFile)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusBadRequest, "Account not found")
	}

	summary, err := usecases.CreateSummaryFromTransactions(transactions, account)
	if err != nil {
		log.Println(err)
		return c.String(http.StatusInternalServerError, "Failed to build the summary")
	}

	err = usecases.SendSummaryByEmail(summary, email)

	return c.String(http.StatusOK, "Email sent!")
}
