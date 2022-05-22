// Package defines usecases for project.
package usecases

import (
	"os"

	"github.com/tlacuilose/transaction-reporter/data/store_reader"
	"github.com/tlacuilose/transaction-reporter/domain/entities"
)

// ReadFromLocalFile reads transactions in a csv file.
func ReadFromLocalFile(filename string) ([]entities.Transaction, error) {
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return make([]entities.Transaction, 0), err
	}

	storeReader := store_reader.New(entities.DecodeFromStrings)
	transactions, err := storeReader.ReadTransactionsFromCsv(f)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
