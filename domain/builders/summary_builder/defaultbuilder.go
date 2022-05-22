// Package defines a summary builder.
package summary_builder

import (
	"errors"

	"github.com/tlacuilose/transaction-reporter/domain/entities"
)

// DefaultSummeryBuilder builds a summary according to SummaryBuilder.
type defaultSummeryBuilder struct {
	transactions           []entities.Transaction
	account                string
	reviewedTransactions   bool
	totalBalance           float64
	numMonthlyTransactions [12]uint
	monthCredit            [12]float64
	monthDebit             [12]float64
	avgCredit              float64
	avgDebit               float64
}

// NewDefaultSummaryBuilder returns a default summary builder.
// The default is getting the summary from a list of entities.Transaction.
func NewDefaultSummaryBuilder(transactions []entities.Transaction) *defaultSummeryBuilder {
	return &defaultSummeryBuilder{
		transactions: transactions,
	}
}

// setAccount sets the name of the account of this report.
func (b *defaultSummeryBuilder) setAccount(account string) {
	b.account = account
}

// getTotalBalance gets the total balance from the list of transactions.
// This function should be called before other methods, besides name.
func (b *defaultSummeryBuilder) getTotalBalance() error {
	for _, t := range b.transactions {
		b.totalBalance += t.Amount

		// Initial range over transactions, sets up for further steps.
		month, err := t.GetMonth()
		if err != nil {
			return err
		}

		monthIndex := month - 1
		b.numMonthlyTransactions[monthIndex]++

		if t.IsCredit() {
			b.monthCredit[monthIndex] += t.Amount
		}

		if t.IsDebit() {
			b.monthDebit[monthIndex] += t.Amount
		}

	}

	b.reviewedTransactions = true
	return nil
}

// getNumMonthlyTransactions gets the number of transactions per month.
// This step has been done in getTotalBalance to prevent further use of for in range.
func (b *defaultSummeryBuilder) getNumMonthlyTransactions() error {
	if !b.reviewedTransactions {
		return errors.New("Default builder requires getTotalBalance to be run first.")
	}
	return nil
}

// getAvgCredit gets the avergae credit per month.
// This step requires getTotalBalance to prevent further use of for in range.
func (b *defaultSummeryBuilder) getAvgCredit() error {
	if !b.reviewedTransactions {
		return errors.New("Default builder requires getTotalBalance to be run first.")
	}

	credit := 0.0
	monthsUsed := 0
	for _, c := range b.monthCredit {
		if c > 0 {
			credit += c
			monthsUsed++
		}
	}

	if monthsUsed > 0 {
		b.avgCredit = credit / float64(monthsUsed)
	}
	return nil
}

// getAvgDebit gets the avergae debit per month.
// This step requires getTotalBalance to prevent further use of for in range.
func (b *defaultSummeryBuilder) getAvgDebit() error {
	if !b.reviewedTransactions {
		return errors.New("Default builder requires getTotalBalance to be run first.")
	}

	debit := 0.0
	monthsUsed := 0
	for _, d := range b.monthDebit {
		if d < 0 {
			debit += d
			monthsUsed++
		}
	}

	if monthsUsed > 0 {
		b.avgDebit = debit / float64(monthsUsed)
	}
	return nil
}

// getSummary returns the summary entitity.
func (b *defaultSummeryBuilder) getSummary() entities.Summary {
	return entities.Summary{
		Account:                b.account,
		TotalBalance:           b.totalBalance,
		NumMonthlyTransactions: b.numMonthlyTransactions,
		AvgCredit:              b.avgCredit,
		AvgDebit:               b.avgDebit,
	}
}
