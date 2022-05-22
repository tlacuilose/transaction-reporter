// Package defines the entities of the project.
package entities

import "testing"

// Test that a summary can be created.
func TestCreateSummary(t *testing.T) {
	var numMonthlyTransactions [12]uint
	numMonthlyTransactions[0] = 2
	numMonthlyTransactions[1] = 4

	summary := Summary{
		Account:                "001",
		TotalBalance:           1333,
		NumMonthlyTransactions: numMonthlyTransactions,
		AvgCredit:              65.90,
		AvgDebit:               -45.04,
	}

	if summary.TotalBalance != 1333 || summary.NumMonthlyTransactions[0] != 2 {
		t.Fatal("Failed to create a summary.")
	}
}
