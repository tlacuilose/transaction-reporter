// Package defines the entities of the project.
package entities

import "fmt"

// Summary defines a summary of account transactions.
type Summary struct {
	Account                string
	TotalBalance           float64
	NumMonthlyTransactions [12]uint
	AvgCredit              float64
	AvgDebit               float64
}

// Print the summary in HTML
func (s Summary) PrintHTML() string {
	return fmt.Sprintf("Summary of: %s", s.Account)
}
