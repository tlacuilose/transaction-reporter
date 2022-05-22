// Package defines a summary builder.
package summary_builder

import "github.com/tlacuilose/transaction-reporter/domain/entities"

// SummaryDirector directs building summaries.
type SummaryDirector struct {
	builder SummaryBuilder
}

// NewSummaryDirector creates a new summary director.
func NewSummaryDirector(b SummaryBuilder) *SummaryDirector {
	return &SummaryDirector{b}
}

// BuildSummary runs all the steps to get a summary.
func (d *SummaryDirector) BuildSummary(account string) (entities.Summary, error) {
	d.builder.setAccount(account)
	if err := d.builder.getTotalBalance(); err != nil {
		return entities.Summary{}, err
	}
	if err := d.builder.getNumMonthlyTransactions(); err != nil {
		return entities.Summary{}, err
	}
	if err := d.builder.getAvgCredit(); err != nil {
		return entities.Summary{}, err
	}
	if err := d.builder.getAvgDebit(); err != nil {
		return entities.Summary{}, err
	}

	return d.builder.getSummary(), nil
}
