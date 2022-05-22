// Package defines a summary builder.
package summary_builder

import "github.com/tlacuilose/transaction-reporter/domain/entities"

// SummaryBuilder defines methods in a summary builder.
type SummaryBuilder interface {
	setAccount(account string)
	getTotalBalance() error
	getNumMonthlyTransactions() error
	getAvgCredit() error
	getAvgDebit() error
	getSummary() entities.Summary
}
