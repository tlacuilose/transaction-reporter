// Package defines a summary builder.
package summary_builder

import (
	"fmt"
	"testing"

	"github.com/tlacuilose/transaction-reporter/domain/entities"
)

// Test that the default builder builds a summary.
func TestDefaultBuilder(t *testing.T) {
	transactions := []entities.Transaction{
		{Id: 0, Date: "2/12", Amount: 100.0},
		{Id: 1, Date: "3/14", Amount: 160.0},
		{Id: 2, Date: "5/11", Amount: -10.0},
		{Id: 3, Date: "6/20", Amount: -20.0},
	}

	expectedAccount := "001"
	expectedBalance := 230.0
	expectedMonthlyTransactions := [12]uint{0, 1, 1, 0, 1, 1, 0, 0, 0, 0, 0, 0}
	expectedAvgCredit := 130.0
	expectedAvgDebit := -15.0

	builder := NewDefaultSummaryBuilder(transactions)
	director := NewSummaryDirector(builder)

	summary, err := director.BuildSummary("001")
	if err != nil {
		t.Fatal("Failed to build a summary")
	}

	fmt.Println(summary)

	if summary.Account != expectedAccount {
		t.Fatal("Failed to set the summery account.")
	}

	if summary.TotalBalance != expectedBalance {
		t.Fatal("Failed to set the summery total balance.")
	}

	if summary.AvgCredit != expectedAvgCredit {
		t.Fatal("Failed to set the summery average credit.")
	}

	if summary.AvgDebit != expectedAvgDebit {
		t.Fatal("Failed to set the summery average debit.")
	}

	for i := 0; i < 12; i++ {
		if summary.NumMonthlyTransactions[i] != expectedMonthlyTransactions[i] {
			t.Fatal("Failed to set the number of transactions per month.")
		}
	}

}
