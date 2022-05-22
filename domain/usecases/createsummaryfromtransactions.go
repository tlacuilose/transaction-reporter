// Package defines usecases for project.
package usecases

import (
	"github.com/tlacuilose/transaction-reporter/domain/builders/summary_builder"
	"github.com/tlacuilose/transaction-reporter/domain/entities"
)

// CreateSummaryFromTransactions creates a summary given transactions and account name.
func CreateSummaryFromTransactions(account string, transactions []entities.Transaction) (entities.Summary, error) {
	builder := summary_builder.NewDefaultSummaryBuilder(transactions)
	director := summary_builder.NewSummaryDirector(builder)

	return director.BuildSummary(account)
}
