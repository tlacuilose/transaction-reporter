// Package defines usecases for project.
package usecases

import (
	"fmt"
	"testing"

	"github.com/tlacuilose/transaction-reporter/domain/entities"
)

// Test that transactions can be read from a test store csv file.
func TestReadFromLocalTestFile(t *testing.T) {
	filename := "../../store/tests/001.csv"

	transactions, err := ReadFromLocalFile(filename)
	if err != nil {
		t.Fatal(err)
	}

	tr := transactions[0]
	if tr.Id != 0 || tr.Date != "7/15" || tr.Amount != 60.5 {
		t.Fatal("Failed to read file correctly")
	}
}

// Test that a summary is produced from transactions.
func TestCreateSummaryFromTransactions(t *testing.T) {
	transactions := []entities.Transaction{
		{Id: 0, Date: "2/12", Amount: 100.0},
		{Id: 1, Date: "3/14", Amount: 160.0},
		{Id: 2, Date: "5/11", Amount: -10.0},
		{Id: 3, Date: "6/20", Amount: -20.0},
		{Id: 4, Date: "2/12", Amount: -100.0},
		{Id: 5, Date: "7/11", Amount: 22.0},
	}

	expectedAccount := "002"
	expectedBalance := 152.0
	expectedMonthlyTransactions := [12]uint{0, 2, 1, 0, 1, 1, 1, 0, 0, 0, 0, 0}
	expectedAvgCredit := 282.0 / 3.0
	expectedAvgDebit := -130.0 / 3.0

	summary, err := CreateSummaryFromTransactions(transactions, "002")
	if err != nil {
		t.Fatal("Failed to build a summary")
	}

	if summary.Account != expectedAccount {
		t.Fatal("Failed to set the summary account.")
	}

	if summary.TotalBalance != expectedBalance {
		t.Fatal("Failed to set the summary total balance.")
	}

	if summary.AvgCredit != expectedAvgCredit {
		t.Fatal("Failed to set the summary average credit.")
	}

	if summary.AvgDebit != expectedAvgDebit {
		t.Fatal("Failed to set the summary average debit.")
	}

	for i := 0; i < 12; i++ {
		if summary.NumMonthlyTransactions[i] != expectedMonthlyTransactions[i] {
			t.Fatal("Failed to set the number of transactions per month.")
		}
	}
}

// Test that a summary can be built from file.
func TestCreateSummaryFromFile(t *testing.T) {
	accountName := "001"
	filename := fmt.Sprintf("../../store/tests/%s.csv", accountName)

	transactions, err := ReadFromLocalFile(filename)
	if err != nil {
		t.Fatal(err)
	}

	expectedAccount := "001"
	expectedBalance := 39.74
	expectedMonthlyTransactions := [12]uint{0, 0, 0, 0, 0, 0, 2, 2, 0, 0, 0, 0}
	expectedAvgCredit := 35.25
	expectedAvgDebit := -15.38

	summary, err := CreateSummaryFromTransactions(transactions, accountName)
	if err != nil {
		t.Fatal("Failed to build a summary")
	}

	if summary.Account != expectedAccount {
		t.Fatal("Failed to set the summary account.")
	}

	if summary.TotalBalance != expectedBalance {
		t.Fatal("Failed to set the summary total balance.")
	}

	if summary.AvgCredit != expectedAvgCredit {
		t.Fatal("Failed to set the summary average credit.")
	}

	if summary.AvgDebit != expectedAvgDebit {
		t.Fatal("Failed to set the summary average debit.")
	}

	for i := 0; i < 12; i++ {
		if summary.NumMonthlyTransactions[i] != expectedMonthlyTransactions[i] {
			t.Fatal("Failed to set the number of transactions per month.")
		}
	}

}
