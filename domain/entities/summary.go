// Package defines the entities of the project.
package entities

// Summary defines a summary of account transactions.
type Summary struct {
	Account                string
	TotalBalance           float64
	NumMonthlyTransactions [12]uint
	AvgCredit              float64
	AvgDebit               float64
}
