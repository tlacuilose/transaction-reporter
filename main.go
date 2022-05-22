package main

import "github.com/tlacuilose/transaction-reporter/presentation/server"

func main() {
	s := server.New(1323)
	s.Start()
}
